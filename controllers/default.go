package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "fmt"

)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:root@(10.64.70.45:3306)/wmoniter?charset=utf8")
}

type MainController struct {
	beego.Controller
}


type RtimeController struct {
	beego.Controller
}
type HistorydataController struct {
	beego.Controller
}
type DataController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *RtimeController) Get() {
	c.TplName = "index.html"
}

func (c *HistorydataController) Get() {
	c.TplName = "historydata.html"
}
func (c *DataController) Get() {
	c.TplName = "data.tpl"
}
