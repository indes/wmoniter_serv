package main

import (
	_ "wmoniter_serv/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

