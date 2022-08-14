package dao

import (
	_ "database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"log"
	entity "telegram-bot/entity"
	"telegram-bot/utils"
)

// init 初始化数据库注册
const query = "select * from bot"

func init() {
	//var bots []entity.Bot
	//count, err := ormer.Raw(query).QueryRows(&bots)
	//log.Printf("查询机器人完成,共:[%d]条数据\n", count)
	//if err != nil {
	//	log.Printf("查询机器人出现异常:[%s]\n", err.Error())
	//}
	//for _, item := range bots {
	//	bytes, _ := json.Marshal(item)
	//	fmt.Printf("%v\n", string(bytes))
	//}
}
func FindBotByUsername(username string) entity.Bot {
	bot := entity.Bot{}
	querySeter := ormer.QueryTable(&bot)
	_ = querySeter.Filter("username", username).One(&bot)
	log.Printf("查询机器人完成,username:[%s],结果:[%v]\n", username, bot)
	return bot
}

func FindBotByUsername1(username string) entity.Bot {
	bot := entity.Bot{Username: username}
	bot.Username = username
	bots := append([]entity.Bot{}, entity.Bot{DelFlag: 1})
	count, err := ormer.Raw(query).QueryRows(&bots)
	log.Printf("查询机器人完成,共:[%d]条数据\n", count)
	if err != nil {
		log.Printf("查询机器人出现异常:[%s]\n", err.Error())
	}
	fmt.Println("---------------")
	for _, item := range bots {
		bytes, _ := json.Marshal(item)
		fmt.Printf("%v\n", string(bytes))
	}
	//fmt.Printf("%v\n", bots)
	//fmt.Printf("%v\n", bot)
	fmt.Println("---------------")
	return entity.Bot{}
}

func SaveBot(bot *entity.Bot) bool {
	temp := FindBotByUsername(bot.Username)
	if utils.IsEmpty(temp.Id) || len(temp.Id) <= 0 {
		UUID, _ := uuid.NewRandom()
		bot.Id = UUID.String()
		length, err := ormer.Insert(bot)
		if err != nil {
			log.Printf("添加机器人出现异常:[%s]\n", err.Error())
			return false
		}
		log.Printf("添加机器人完成,添加记录:[%d]条\n", length)
		return length >= 0
	} else {
		temp.GroupName = bot.GroupName
		temp.Token = bot.Token
		temp.ChatId = bot.ChatId
		temp.BotName = bot.BotName
		temp.CallbackUrl = bot.CallbackUrl
		length, err := ormer.Update(temp)
		if err != nil {
			log.Printf("更新机器人出现异常:[%s]\n", err.Error())
			return false
		}
		log.Printf("更新机器人完成,添加记录:[%d]条\n", length)
		return length >= 0
	}
}
