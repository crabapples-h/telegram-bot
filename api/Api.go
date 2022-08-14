package api

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"telegram-bot/dao"
	entity "telegram-bot/entity"
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
	log.Printf("收到回调请求->参数->[%s]\n", bodyStr)
	callBackData := form.CallBackData{}
	err := json.Unmarshal([]byte(bodyStr), &callBackData)
	if err != nil {
		log.Printf("json反序列化出现异常:[%s]", err.Error())
	}
	doSomething(callBackData)
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
		if !utils.IsEmpty(newChatMember) {
			var isBot = newChatMember.IsBot
			if isBot {
				log.Println("开始更新机器人的chatId")
				var chat = message.Chat
				var chatId = chat.Id
				var botUserName = newChatMember.Username
				var groupName = chat.Title
				var bot = service.FindByUsername(botUserName)
				bot.ChatId = chatId
				bot.GroupName = groupName
				if service.SaveBot(&bot) {
					log.Println("更新机器人的chatId完成")
				} else {
					log.Println("更新机器人的chatId失败")
				}
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
			if !utils.IsEmpty(replyToMessage) {
				log.Printf("消息类型为回复,内容:[%v]\n", replyToMessage)
				cacheMessageId := replyToMessage.MessageId
				cacheMessageStr, _ := dao.Redis.Get(strconv.FormatInt(cacheMessageId, 10)).Result()
				if len(cacheMessageStr) <= 0 {
					service.SendMessage("操作无效或已超时，请重新开始", token, chatId)
					return
				}
				cacheMessage := lib.Message{}
				err := json.Unmarshal([]byte(cacheMessageStr), &cacheMessage)
				if err != nil {
					fmt.Println(err.Error())
				}
				cacheMessageText := cacheMessage.Text
				if "请回复此消息机器人创建时的用户名及token,格式为：xxxxx_bot##token" == (cacheMessageText) {
					split := strings.Split(messageContent, "##")
					if len(split) != 2 {
						service.SendMessage("回复格式错误，请重试", token, chatId)
						return
					}
					botUsername := split[0]
					botToken := split[1]
					log.Printf("创建机器人,username:[%s],token:[%s]\n", botUsername, botToken)
					bot := entity.Bot{
						Token:    botToken,
						Username: botUsername,
					}
					log.Println("开始创建机器人")
					if service.SaveBot(&bot) {
						log.Println("创建机器人完成")
					} else {
						log.Println("创建机器人失败")
					}
					dao.Redis.Del(strconv.FormatInt(cacheMessageId, 10))
					service.SendMessage("很好，请将该机器人加入需要发送消息的群组，并设为管理员", token, chatId)
				}
			}
			log.Printf("回复的内容:[%v]\n", replyToMessage)
		}
	}
}
