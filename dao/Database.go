package dao

import (
	_ "database/sql"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"telegram-bot/conf"
	entity "telegram-bot/entity"
)

var ormer orm.Ormer

func init() {
	log.Println("开始初始化数据库连接")
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//注册模型类似jpa关联
	models := []interface{}{
		new(entity.Bot),
		new(entity.Command),
	}
	orm.RegisterModel(models...)
	dbUrl := fmt.Sprintf("%s:%s@(%s:%d)/%s?%s", conf.DbUsername, conf.DbPassword, conf.DbHost, conf.DbPort, conf.DbName, "charset=utf8")
	orm.RegisterDataBase(conf.DbAlias, conf.DbDriveName, dbUrl)
	orm.SetMaxIdleConns(conf.DbAlias, conf.MaxIdle)
	orm.SetMaxOpenConns(conf.DbAlias, conf.MaxConn)
	// 创建一个全局Ormer对象
	ormer = orm.NewOrm()
	log.Println("初始化数据库连接完成")
}
