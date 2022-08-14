package dao

import (
	_ "database/sql"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	entity "telegram-bot/entity"
)

var ormer orm.Ormer

const (
	dbDriveName = "mysql"
	dbAlias     = "default"
	dbName      = "telegram"
	dbHost      = "127.0.0.1"
	dbPort      = 3306
	dbUsername  = "root"
	dbPassword  = ""
	maxIdle     = 30 //最大空闲连接
	maxConn     = 30 //最大数据库连接

)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//注册模型类似jpa关联
	models := []interface{}{
		new(entity.Bot),
		new(entity.Command),
	}
	orm.RegisterModel(models...)
	dbUrl := fmt.Sprintf("%s:%s@(%s:%d)/%s?%s", dbUsername, dbPassword, dbHost, dbPort, dbName, "charset=utf8")
	orm.RegisterDataBase(dbAlias, dbDriveName, dbUrl)
	orm.SetMaxIdleConns(dbAlias, maxIdle)
	orm.SetMaxOpenConns(dbAlias, maxConn)
	// 创建一个全局Ormer对象
	ormer = orm.NewOrm()
}
