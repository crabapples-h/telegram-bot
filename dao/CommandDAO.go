package dao

import (
	_ "database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	entity "telegram-bot/entity"
)

func FindCommand(commandStr string) entity.Command {
	command := entity.Command{}
	querySeter := ormer.QueryTable(&command)
	_ = querySeter.Filter("command", commandStr).One(&command)
	log.Printf("查询命令完成,command:[%s],结果:[%v]\n", commandStr, command)
	return command
}
