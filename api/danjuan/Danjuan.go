package danjuan

import (
	"fmt"
	"io/ioutil"
	"net/http"

	jsoniter "github.com/json-iterator/go"
)

const host string = "https://danjuanapp.com"

type IndexResponse struct {
	IndexData  IndexData `json:"data"`
	ResultCode int       `json:"result_code"`
}
type Items struct {
	ID            int     `json:"id"`
	IndexCode     string  `json:"index_code"`
	Name          string  `json:"name"`
	Ttype         string  `json:"ttype"`
	Pe            float64 `json:"pe"`
	Pb            float64 `json:"pb"`
	PePercentile  float64 `json:"pe_percentile"`
	PbPercentile  float64 `json:"pb_percentile"`
	Roe           float64 `json:"roe"`
	Yeild         float64 `json:"yeild"`
	Ts            int64   `json:"ts"`
	EvaType       string  `json:"eva_type"`
	EvaTypeInt    int     `json:"eva_type_int"`
	URL           string  `json:"url"`
	BondYeild     float64 `json:"bond_yeild"`
	BeginAt       int64   `json:"begin_at"`
	CreatedAt     int64   `json:"created_at"`
	UpdatedAt     int64   `json:"updated_at"`
	Peg           float64 `json:"peg,omitempty"`
	PbFlag        bool    `json:"pb_flag"`
	Date          string  `json:"date"`
	PbOverHistory float64 `json:"pb_over_history"`
	PeOverHistory float64 `json:"pe_over_history"`
}
type IndexData struct {
	Items       []Items `json:"items"`
	CurrentPage int     `json:"current_page"`
	Size        int     `json:"size"`
	TotalItems  int     `json:"total_items"`
	TotalPages  int     `json:"total_pages"`
}

//获取指数信息
func GetIndex() []Items {
	path := "/djapi/index_eva/dj"
	var url = fmt.Sprintf("%s%s", host, path)
	fmt.Println("url:" + url)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println("body:", string(body))
	var indexResponse IndexResponse // := new(IndexResponse)
	jsoniter.Unmarshal(body, &indexResponse)
	return indexResponse.IndexData.Items
}
