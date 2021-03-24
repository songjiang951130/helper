package xueqiu

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

const host string = "https://stock.xueqiu.com"

var header http.Header
var mu sync.Mutex

//获取某只股票前十大持有人
func GetTopHolder(symbol string) {
	path := "/v5/stock/f10/cn/top_holders.json"
	var url = fmt.Sprintf("%s%s?symbol=%s&circula=1&count=200", host, path, symbol)
	fmt.Println("url:" + url)
	client := &http.Client{}
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}
	request.Header = GetHeader()
	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("body:", string(body))
}

func GetHeader() http.Header {
	if header == nil {
		mu.Lock()
		defer mu.Unlock()
		if header == nil {
			cookie := GetCookie()
			header = map[string][]string{
				"Cookie":    {cookie},
				"Origin":    {host},
				"sec-ch-ua": {host},
			}
		}
	}
	return header
}
