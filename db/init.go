package db

import (
	"SoupSoul/models"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
	"sync"
)

var once sync.Once
var dao *orm.Ormer

func init() {
	Instance()
}

func Instance() orm.Ormer {
	once.Do(func() {
		dbUser, _ := web.AppConfig.String("mysqluser")
		dbPass, _ := web.AppConfig.String("mysqlpass")
		dbAddr, _ := web.AppConfig.String("mysqladdress")
		dbPort, _ := web.AppConfig.String("mysqlport")
		dbName, _ := web.AppConfig.String("mysqldb")

		dbSrouce := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbUser, dbPass, dbAddr, dbPort, dbName)
		orm.RegisterDriver("mysql", orm.DRMySQL)
		orm.RegisterDataBase("default", "mysql", dbSrouce)
		orm.SetMaxIdleConns("default", 30)
		orm.SetMaxOpenConns("default", 30)

		orm.RegisterModelWithPrefix("ai_", new(models.Soul))

		// 自创建表
		orm.RunSyncdb("default", false, true)

		d := orm.NewOrm()
		dao = &d
	})
	return *dao
}
