package api

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"telegram-bot/dao"
	"telegram-bot/form"
	"telegram-bot/lib"
	_ "telegram-bot/lib"
	"telegram-bot/service"
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
	doSomething(callBackData)
	log.Printf("收到回调请求->参数->[%s]\n", bodyStr)
}

func doSomething(form form.CallBackData) {
	var callbackQuery = form.CallbackQuery
	if !utils.IsEmpty(callbackQuery) {
		var callbackQueryData = callbackQuery.Data
		if "random:message" == callbackQueryData {
			log.Println("触发随机推送")
			var chat = callbackQuery.Message.Chat
			var chatId = chat.Id
			var botUsername = callbackQuery.Message.From.Username
			service.RandomSend(botUsername, chatId)
			return
		}
	}
	var message = form.Message
	if !utils.IsEmpty(message) {
		var newChatMember = message.NewChatMember
		log.Printf("%v\n", newChatMember)
		log.Println(utils.IsEmpty(newChatMember))
		if !utils.IsEmpty(newChatMember) {
			var isBot = newChatMember.IsBot
			if isBot {
				var chat = message.Chat
				var chatId = chat.Id
				var botUserName = newChatMember.Username
				var groupName = chat.Title
				var bot = service.FindByUsername(botUserName)
				bot.ChatId = chatId
				bot.GroupName = groupName
				service.SaveBot(bot)
				return
			}
		}
		enumMessage(message)
	}

}
func enumMessage(message lib.Message) {
	log.Println("开始查找消息命令")
	token := "5448913855:AAGMPRqmMswqV3Il_LtpAk5fEJg45uAMjvc"
	chat := message.Chat
	if !utils.IsEmpty(chat) {
		if "private" == chat.Type {
			messageContent := message.Text
			chatId := chat.Id
			command := dao.FindCommand(messageContent)
			if len(command.Id) > 0 {
				log.Printf("找到命令:[%v]\n", command)
				replyMessage := command.ReplyMessage
				messageResult := service.SendMessage(replyMessage, token, chatId)
				log.Printf("回复的内容:[%v]\n", messageResult)
				return
			}
			replyToMessage := message.ReplyToMessage
			log.Printf("回复的内容:[%v]\n", replyToMessage)
		}
	}
}
