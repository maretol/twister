package logic

import (
	"io/ioutil"
	"log"
	"net/http"
)

// Booster は path とそれに対応するHTTPアクセスの結果の構造体
type Booster struct {
	path         string
	urlList      []localUrls
	ResultMap    map[string]Fuel
	HeaderIgnore []string
	Checker      struct {
		StatusCode bool
		Header     bool
		Body       bool
	}
}

// Fuel はHTTPアクセスの比較に使われる要素をひとまとめにした構造体
type Fuel struct {
	StatusCode int
	Header     map[string][]string
	Body       string
}

type localUrls struct {
	fullURL string
	tag     string
}

// GetURL は cmd.URL を受け取るために使う構造体。変換して使って
type GetURL struct {
	URL string
	Tag string
}

// SetFullURL は各URLにPATHをつなげたURLを作ってそれを Booster 構造体にセットします
func (b *Booster) SetFullURL(path string, urls []GetURL) {
	b.path = path
	b.urlList = make([]localUrls, len(urls))
	b.ResultMap = make(map[string]Fuel, len(urls))
	for num, url := range urls {
		b.urlList[num].fullURL = url.URL + path
		b.urlList[num].tag = url.Tag
	}
}

// AllAccess はSetFullURLで設定したURLにアクセスします。
func (b *Booster) AllAccess() {
	for num := range b.urlList {
		url := b.urlList[num].fullURL
		tag := b.urlList[num].tag
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
				Header:     res.Header,
				StatusCode: res.StatusCode,
				Body:       string(body),
			}
		}()
		b.ResultMap[tag] = fuel
	}
}
