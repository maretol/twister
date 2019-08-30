package logic

import (
	"io/ioutil"
	"log"
	"net/http"
)

// Booster は path とそれに対応するHTTPアクセスの結果の構造体
type Booster struct {
	path      string
	urlList   []string
	ResultMap map[string]Fuel
}

// Fuel はHTTPアクセスの比較に使われる要素をひとまとめにした構造体
type Fuel struct {
	StatusCode int
	Header     string // Header は Key/Value にするかも
	Body       string
}

// SetFullURL は各URLにPATHをつなげたURLを作ってそれを Booster 構造体にセットします
func (b *Booster) SetFullURL(path string, urls []string) {
	b.path = path
	b.urlList = urls
	b.ResultMap = make(map[string]Fuel, len(urls))
	for num := range urls {
		url := urls[num]
		fullURL := url + path
		b.urlList = append(b.urlList, fullURL)
	}
}

// AllAccess はSetFullURLで設定したURLにアクセスします。
func (b *Booster) AllAccess() {
	for num := range b.urlList {
		url := b.urlList[num]
		// channel := make(chan Fuel) // 最終的には並列にしたい
		var fuel Fuel
		func() {
			res, err := http.Get(url)
			if err != nil {
				log.Fatal(err)
			}
			body, _ := ioutil.ReadAll(res.Body)
			res.Body.Close()
			fuel = Fuel{
				StatusCode: res.StatusCode,
				Body:       string(body),
			}
		}()
		b.ResultMap[url] = fuel
	}
}

// ShowDiff は結果の差分を表示するメソッドです
func (b *Booster) ShowDiff() {

}
