package main

import (
	_ "SoupSoul/db"
	_ "SoupSoul/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	defer beego.Run()

	// 初始化逻辑

}
