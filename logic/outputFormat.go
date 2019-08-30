package logic

import (
	"fmt"
	"strings"

	"github.com/sergi/go-diff/diffmatchpatch"
)

// Output に結果を渡してくれればあとは処理をします
func Output(result Booster, statusCodeCheck bool, headerCheck bool, bodyCheck bool) {
	statusCodeMatched := statusCodeCheck && isStatusCodeMatched(result.ResultMap)
	headerMatched := headerCheck && isHeaderMatched(result.ResultMap) // Headerの一致処理は今は実装しない
	bodyMatched := bodyCheck && isBodyMatched(result.ResultMap)
	matchResult := setResult(statusCodeMatched, bodyMatched) // headerMatched はそのまま追加すればいい
	// 結果部分1（概要。一致したかしていないか）
	fmt.Printf(resultFormat1, result.path, matchResult)
	// 結果部分2（ステータスコード。一致していなかったときだけ表示
	if statusCodeMatched {
		fmt.Printf(resultFormat2, strings.Join(getStatusCodeList(result.ResultMap), ","))
	}
	// 結果部分3（ヘッダー、一致していなかったときだけ表示。現状はすべて表示
	if headerMatched {
		urls := getKeysFromMap(result.ResultMap)
		for _, url := range urls {
			fmt.Printf(resultFormat3_2, url, "") // 後ろの空文字はHeaderの差分
		}
	}
	// 結果部分4（Body部分。一致していなかったときだけ表示
	// ここは差分を出したい
	if bodyMatched {
		bodyDiff := getBodyDiff(result.ResultMap)
		fmt.Printf(resultFormat4_1)
		fmt.Printf(resultFormat4_2, bodyDiff)
	}
}

// これ Config のステータスが合致していないケースを入れると若干ややこしいのでそこを踏まえてちょっと書き直す
func setResult(matchers ...bool) string {
	for num := range matchers {
		if !matchers[num] { // どれか1つでも false なら一致していないので
			return "Not matched"
		}
	}
	return "Matched"
}

func isStatusCodeMatched(ResultMap map[string]Fuel) bool {
	previousStatusCode := 0
	for _, response := range ResultMap {
		statusCode := response.StatusCode
		if statusCode != previousStatusCode && previousStatusCode != 0 {
			return false
		}
		previousStatusCode = response.StatusCode
	}
	return true
}

// header の各要素を比較する。全部比較したいけど結果をどうするか（タイムスタンプとかほぼ100％異なるし
func isHeaderMatched(ResultMap map[string]Fuel) bool {
	// 今は実施しない
	// dmp := diffmatchpatch.New()
	// var previousHeader http.Header
	// for _, response := range ResultMap {
	// 	if previousHeader == nil {
	// 		previousHeader = response.Header
	// 		continue
	// 	}
	// 	for headerKey, _ := range response.Header {
	// 		dmp.DiffLinesToChars(response.Header.Get((headerKey)), previousHeader.Get((headerKey)))
	// 	}
	// }
	return false
}

// 一致しなかったテキストのペアは別に出す
// TODO: ContentType が JSON や XML のときには内容で比較できるようにしたい
func isBodyMatched(ResultMap map[string]Fuel) bool {
	urls := getKeysFromMap(ResultMap)
	url1 := urls[0]
	url2 := urls[1]

	body1 := ResultMap[url1].Body
	body2 := ResultMap[url2].Body

	return body1 != body2
}

func getBodyDiff(ResultMap map[string]Fuel) []string {
	dmp := diffmatchpatch.New()

	urls := getKeysFromMap(ResultMap)
	url1 := urls[0]
	url2 := urls[1]
	body1 := ResultMap[url1].Body
	body2 := ResultMap[url2].Body
	b1, b2, dump := dmp.DiffLinesToChars(body1, body2)
	diffs := dmp.DiffMain(b1, b2, false)
	dmp.DiffCharsToLines(diffs, dump)

	var result []string
	for _, diff := range diffs {
		if diff.Type != diffmatchpatch.DiffEqual {
			// ここで同じじゃない行がdiff経由で扱える
			result = append(result, diff.Text)
		}
	}
	return result
}

func getStatusCodeList(ResultMap map[string]Fuel) []string {
	var tmp []string
	for _, response := range ResultMap {
		tmp = append(tmp, (string)(response.StatusCode))
	}
	return tmp
}

// ここでしか使わないので型は固定している（本当は interface{} で受けたかったけどできなかった
func getKeysFromMap(fromMap map[string]Fuel) []string {
	keys := make([]string, 0, len(fromMap))
	for k := range fromMap {
		keys = append(keys, k)
	}
	return keys
}

const resultFormat1 = `
---------------------------------------------------------------------------------------
PATH   : %s
RESULT : %s
`

const resultFormat2 = `
STATUS CODE : %s
`

const resultFormat3_1 = `
HEADER
`
const resultFormat3_2 = `
URL : %s
	%s
`

const resultFormat4_1 = `
BODY
`
const resultFormat4_2 = `
	Line : - (今は出せない)
	Diff :
%s
`
