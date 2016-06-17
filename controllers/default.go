package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"

	_ "fmt"
	"fmt"
	"wmoniter_serv/models"
)


func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:root@/wmoniter?charset=utf8")
}

type MainController struct {
	beego.Controller
}
type RtimeController struct {
	beego.Controller
}

type DatabaseController struct {
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
func (c *DatabaseController) Get() {
	o := orm.NewOrm()
	o.Using("default")
	orm.RunSyncdb("default", false, true)
	c.Ctx.WriteString("123\n")
	//user := new(Data)
	//user.Id=1213
	//user.Date=time.Now()
	//user.Value = 213321213
	//o.Insert(user)
	u :=models.Data{Id: 6}
	err := o.Read(&u)
	fmt.Printf("ERR: %v\n", err)
	c.Ctx.WriteString(u.Date.Format("2006-01-02 15:04:05:00"))

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

