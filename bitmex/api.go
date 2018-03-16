package bitmex

import (
	"fmt"
	"net/url"
	"strings"
	"io/ioutil"
	"encoding/json"
	"net/http"
	"log"
)

const (
	UA      = "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 Safari/537.36"
	BaseURL = "https://testnet.bitmex.com/api/v1"
)

type Quote struct {
	Timestamp string  `json:"timestamp"`
	Symbol    string  `json:"symbol"`
	BidSize   int     `json:"bidSize"`
	BidPrice  float32 `json:"bidPrice"`
	AskPrice  float32 `json:"askPrice"`
	AskSize   int     `json:"askSize"`
}

func queryString(opt map[string]string) string {
	var pairStr []string
	for k, v := range opt {
		pairStr = append(pairStr, fmt.Sprintf("%s=%s", k, url.QueryEscape(v)))
	}
	log.Print(pairStr)
	return strings.Join(pairStr, "&")
}


// Quotes 行情数据
func Quotes(opt map[string]string) (q []Quote, err error){

	// 接口地址
	endPoint := "/quote"

	// 参数
	qs := queryString(opt)
	log.Print(qs)

	//  拼URL
	url := fmt.Sprintf("%s%s?%s", BaseURL, endPoint, qs)

	log.Print(url)

	// 初始化
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)

	// 设置UA
	req.Header.Set("User-Agent", UA)
	req.Header.Set("Upgrade-Insecure-Requests", "1")

	// 访问
	resp, err := client.Do(req)
	if err != nil {
		return
	}

	// 检查状态码
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("access failed")
		return
	}

	// 获取内容
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(bodyBytes, &q)

	return
}
