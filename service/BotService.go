package service

import entity "telegram-bot/entity"
import "telegram-bot/dao"

func FindByUsername(username string) entity.Bot {
	return dao.FindBotByUsername(username)
}

func SaveBot(bot entity.Bot) bool {
	return dao.SaveBot(bot)
}
