package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"telegram-bot/dao"
	"telegram-bot/entity"
	"telegram-bot/form"
	"telegram-bot/utils"
	"time"
)

func RandomSend(username string, chatId int64) {

}

const ApiUrl = "https://api.telegram.org/bot%s/%s"

func SendMessage(message, token string, chatId int64) form.MessageResult {
	method := "sendMessage"
	data := make(map[string]interface{})
	data["chat_id"] = chatId
	data["text"] = message
	byteData := request(method, token, data)
	result := form.MessageResult{}
	err := json.Unmarshal(byteData, &result)
	if err != nil {
		log.Panicf("消息反序列化失败:[%s]\n", err.Error())
	}
	sendMessage := result.Result
	messageId := sendMessage.MessageId
	messageJson, _ := json.Marshal(sendMessage)
	dao.Redis.Set(strconv.FormatInt(messageId, 10), messageJson, time.Minute*10)
	return result
}

func UpdateWebHook(bot *eneity.Bot) {
	callbackUrl := bot.CallbackUrl
	if len(callbackUrl) <= 0 {
		return
	}
	token := bot.Token
	removeMethod := "deleteWebhook"
	removeData := request(removeMethod, token, nil)
	log.Printf("删除webhook:[%s]", string(removeData))
	setMethod := "setWebhook"
	data := make(map[string]interface{})
	data["url"] = callbackUrl
	setData := request(setMethod, token, data)
	log.Printf("设置webhook:[%s]", string(setData))
}

func request(method, token string, data map[string]interface{}) []byte {
	apiUrl := fmt.Sprintf(ApiUrl, token, method)
	client := http.Client{Timeout: time.Second * 30}
	utils.SetHttpProxy(&client)
	dataJson, _ := json.Marshal(data)
	req, err := http.NewRequest("POST", apiUrl, bytes.NewReader(dataJson))
	if err != nil {
		log.Println("网络请求失败", err.Error())
	}
	req.Header.Add("Content-Type", "application/json")
	response, err := client.Do(req)
	if err != nil {
		log.Panicf("网络请求出现异常:[%s]\n", err.Error())
	}
	body := response.Body
	all, err := ioutil.ReadAll(body)
	return all
}
