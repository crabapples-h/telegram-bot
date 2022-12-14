package utils

import (
	"log"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

const HttpProxyUrl = "http://127.0.0.1:10800"

func Unicode2Zh(raw []byte) string {
	log.Printf("开始转换字符串[unicode]->[中文]")
	str, err := strconv.Unquote(strings.Replace(strconv.Quote(string(raw)), `\\u`, `\u`, -1))
	if err != nil {
		log.Printf("转换异常:[%s]", err.Error())
		return ""
	}
	log.Printf("字符串转换完成[unicode]->[中文]")
	return str
}

func IsEmpty(obj interface{}) bool {
	return reflect.DeepEqual(obj, nil)
}
func SetHttpProxy(client *http.Client) {
	ProxyURL, _ := url.Parse(HttpProxyUrl)
	client = &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(ProxyURL),
		},
	}
}
