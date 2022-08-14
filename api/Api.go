package api

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"telegram-bot/form"
	"telegram-bot/lib"
	"telegram-bot/utils"
)

func CallBack(context *gin.Context) {
	fmt.Println(context)
	bodyByte, _ := ioutil.ReadAll(context.Request.Body)
	bodyStr := utils.Unicode2Zh(bodyByte)
	callBackData := form.CallBackData{}
	err := json.Unmarshal([]byte(bodyStr), &callBackData)
	if err != nil {
		log.Printf("json反序列化出现异常:[%s]", err.Error())
	}
	//marshal, err := json.Marshal(&callBackData)
	fmt.Println(bodyStr)
	aa(callBackData)
	log.Printf("收到回调请求->参数->[%s]\n", bodyStr)
}

func aa(form form.CallBackData) {
	var callbackQuery = form.CallbackQuery
	//lib.Message = form.Message
	fmt.Printf("%v\n", callbackQuery)
	fmt.Println(utils.IsEmpty(callbackQuery, lib.CallbackQuery{}))
}
func enumMessage() {

}
