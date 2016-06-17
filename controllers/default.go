package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"

	_ "fmt"
	"math/rand"
	"wmoniter_serv/models"
	"time"
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
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	o := orm.NewOrm()
	o.Using("default")
	orm.RunSyncdb("default", false, true)
	c.Ctx.WriteString("123\n")
	for i := 1; i <= 1000; i++ {
		user := new(models.Data)
		user.Id = 1213
		user.Date = time.Now()
		user.Value = r.Intn(10)
		o.Insert(user)
	}

	//u :=models.Data{Id: 6}
	//err := o.Read(&u)
	//fmt.Printf("ERR: %v\n", err)
	//c.Ctx.WriteString(u.Date.Format("2006-01-02 15:04:05:00"))

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

