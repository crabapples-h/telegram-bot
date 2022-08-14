package utils

import (
	"log"
	"reflect"
	"strconv"
	"strings"
)

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
