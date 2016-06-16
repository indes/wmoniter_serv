package controllers

import (
	"github.com/astaxie/beego"
)

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
	//c.Ctx.ResponseWriter.Header("Content-type: text/javascript;charset=UTF-8")
	c.TplName = "data.tpl"
}