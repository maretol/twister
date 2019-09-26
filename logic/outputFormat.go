package logic

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/sergi/go-diff/diffmatchpatch"
)

// Output に結果を渡してくれればあとは処理をします
func Output(result Booster) {
	statusCodeMatched := result.Checker.StatusCode && needsStatusCodeMatched(result.ResultMap)
	headerMatched := result.Checker.Header && needsHeaderMatched(result.ResultMap, result.HeaderIgnore)
	bodyMatched := result.Checker.Body && needsBodyMatched(result.ResultMap)

	showResult(result, statusCodeMatched, headerMatched, bodyMatched)

}

func showResult(result Booster, statusCodeMatched, headerMatched, bodyMatched bool) {
	matchResult := setResult(statusCodeMatched, bodyMatched) // headerMatched はそのまま追加すればいい

	// 結果部分1（概要。一致したかしていないか）
	fmt.Printf(resultFormat1, matchResult, result.path)

	// 結果部分2（ステータスコード。一致していなかったときだけ表示
	if statusCodeMatched {
		fmt.Printf(resultFormat2, strings.Join(getStatusCodeList(result.ResultMap), "\n    "))
	}

	// 結果部分3（ヘッダー、一致していなかったときだけ表示。現状はすべて表示
	if headerMatched {
		// 双方の header のキーを一つの配列に入れて重複削除しそれぞれを比較する感じで
		fmt.Printf(resultFormat3_1)
		urls := GetKeysFromFuelMap(result.ResultMap)
		url1 := urls[0]
		url2 := urls[1]
		header1 := result.ResultMap[url1].Header
		header2 := result.ResultMap[url2].Header
		baseKey := GetKeysFromStringSliceMap(header1)
		key2 := GetKeysFromStringSliceMap(header2)
		baseKey = append(baseKey, key2...)
		singleKeys := make([]string, 0, len(baseKey))
		encountered := map[string]bool{}
		for _, v := range result.HeaderIgnore {
			encountered[v] = true
		}
		for _, v := range baseKey {
			if !encountered[v] {
				encountered[v] = true
				singleKeys = append(singleKeys, v)
			}
		}
		for _, key := range singleKeys {
			data1 := header1[key]
			data2 := header2[key]
			for _, value1 := range data1 {
				match := false
				for _, value2 := range data2 {
					match = match || value1 == value2
				}
				if !match {
					fmt.Printf(resultFormat3_2, url1, key, data1)
					fmt.Printf(resultFormat3_2, url2, key, data2)
					fmt.Println()
				}
			}
		}
		fmt.Println()
	}

	// 結果部分4（Body部分。一致していなかったときだけ表示
	// ここは差分を出したい
	if bodyMatched {
		bodyDiff := getBodyDiff(result.ResultMap)
		fmt.Printf(resultFormat4_1)
		fmt.Printf(resultFormat4_2)
		for num, diff := range bodyDiff {
			num2 := num
			for num2 >= len(result.urlList) {
				num2 = num2 - len(result.urlList)
			}
			fmt.Printf(resultFormat4_3, result.urlList[num2].tag, diff)
		}
	}
}

// これ Config のステータスが合致していないケースを入れると若干ややこしいのでそこを踏まえてちょっと書き直す
func setResult(matchers ...bool) string {
	for num := range matchers {
		if matchers[num] { // どれか1つでも true なら表示するものがあるので
			return "FAIL"
		}
	}
	return "SUCCESS"
}

func needsStatusCodeMatched(ResultMap map[string]Fuel) bool {
	previousStatusCode := 0
	for _, response := range ResultMap {
		statusCode := response.StatusCode
		if statusCode != previousStatusCode && previousStatusCode != 0 {
			return true
		}
		previousStatusCode = response.StatusCode
	}
	return false
}

// header の各要素を比較する。headerIgnore を回避する必要あり
func needsHeaderMatched(ResultMap map[string]Fuel, headerIgnore []string) bool {
	urls := GetKeysFromFuelMap(ResultMap)
	url1 := urls[0]
	url2 := urls[1]
	header1 := ResultMap[url1].Header
	header2 := ResultMap[url2].Header
	// ignore で無視したいHeaderのキーを上書きする
	// これで「片方にはない」「両方あるがそれぞれ違う」「両方ともない」すべてがカバーできる
	for _, v := range headerIgnore {
		header1[v] = nil
		header2[v] = nil
	}

	// 要素数がすでに違うので false
	if len(header1) != len(header2) {
		return true
	}

	for k, v1 := range header1 {
		v2 := header2[k]
		if len(v1) != len(v2) {
			return true
		}
		for _, str1 := range v1 {
			for _, str2 := range v2 {
				if str1 != str2 {
					return true
				}
			}
		}
	}
	return false
}

// 一致しなかったテキストのペアは別に出す
// TODO: ContentType が JSON や XML のときには内容で比較できるようにしたい
func needsBodyMatched(ResultMap map[string]Fuel) bool {
	urls := GetKeysFromFuelMap(ResultMap)
	url1 := urls[0]
	url2 := urls[1]

	body1 := ResultMap[url1].Body
	body2 := ResultMap[url2].Body

	return body1 != body2
}

func getBodyDiff(ResultMap map[string]Fuel) []string {
	dmp := diffmatchpatch.New()

	urls := GetKeysFromFuelMap(ResultMap)
	url1 := urls[0]
	url2 := urls[1]
	body1 := ResultMap[url1].Body
	body2 := ResultMap[url2].Body
	b1, b2, dump := dmp.DiffLinesToChars(body1, body2)
	diffs := dmp.DiffMain(b1, b2, false)
	diffResult := dmp.DiffCharsToLines(diffs, dump)

	var result []string
	for _, diff := range diffResult {
		if diff.Type != diffmatchpatch.DiffEqual {
			// ここで同じじゃない行がdiff経由で扱える
			result = append(result, strings.TrimRight(diff.Text, "\n"))
		}
	}
	return result
}

func getStatusCodeList(ResultMap map[string]Fuel) []string {
	var tmp []string
	for key, response := range ResultMap {
		tmp = append(tmp, key+":"+strconv.Itoa(response.StatusCode))
	}
	return tmp
}

const resultFormat1 = `
======================================================================================
RESULT : %s
PATH   : %s
`

const resultFormat2 = `
STATUS CODE :
    %s
`

const resultFormat3_1 = `
HEADER
`
const resultFormat3_2 = `URL : %v
%v : %v
`

const resultFormat4_1 = `
BODY`
const resultFormat4_2 = `
Line : - (今は出せない)
Diff :
`
const resultFormat4_3 = `Tag(or URL) : %s
--------------------------------------------------------------------------------------
%s
--------------------------------------------------------------------------------------

`
