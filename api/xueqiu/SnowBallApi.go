package xueqiu

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const host string = "https://stock.xueqiu.com"

func GetTopHolder(symbol string) {
	client := &http.Client{}

	cookie := GetCookie()
	fmt.Println(client, cookie)

	path := "/v5/stock/f10/cn/top_holders.json"
	var url = fmt.Sprintf("%s%s?symbol=%s&circula=1&count=200", host, path, symbol)
	fmt.Println("url:" + url)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	//延迟关闭
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println("body:", string(body))

}
