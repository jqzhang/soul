package controllers

import (
	"SoupSoul/db"
	"SoupSoul/models"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Content"] = getRandContent()
	c.TplName = "index.html"
}

func getRandContent() string {
	var soul models.Soul
	err := db.Instance().Raw("select * from ai_soul order by rand() limit 1").QueryRow(&soul)
	if err != nil {
		fmt.Println("~~~ e ", err.Error())
		return "精神和神经能差多少？"
	} else {
		return soul.Content
	}
}
