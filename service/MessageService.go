package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"telegram-bot/form"
	"time"
)

func RandomSend(username string, chatId int64) {

}

const ApiUrl = "https://api.telegram.org/bot%s/%s"
const HttpProxyUrl = "http://127.0.0.1:10800"

func SendMessage(message, token string, chatId int64) form.MessageResult {
	method := "sendMessage"
	apiUrl := fmt.Sprintf(ApiUrl, token, method)
	client := http.Client{Timeout: time.Second * 30}
	SetHttpProxy(&client)
	data := make(map[string]interface{})
	data["chat_id"] = chatId
	data["text"] = message
	marshal, _ := json.Marshal(data)
	request, err := http.NewRequest("POST", apiUrl, bytes.NewReader(marshal))
	if err != nil {
		log.Println("消息发送失败", err.Error())
	}
	request.Header.Add("Content-Type", "application/json")
	response, err := client.Do(request)
	if err != nil {
		log.Panicf("发送消息出现异常:[%s]\n", err.Error())
	}
	body := response.Body
	all, err := ioutil.ReadAll(body)
	log.Println(string(all))
	result := form.MessageResult{}
	err = json.Unmarshal(all, &result)
	if err != nil {
		log.Panicf("消息反序列化失败:[%s]\n", err.Error())
	}
	return result
}

func SetHttpProxy(client *http.Client) {
	ProxyURL, _ := url.Parse(HttpProxyUrl)
	client = &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(ProxyURL),
		},
	}
}
