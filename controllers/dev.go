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

type DevController struct {
	beego.Controller
}

func (c *DevController) GData() {
	o := orm.NewOrm()
	//生成随机数据
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	o.Using("default")
	orm.RunSyncdb("default", false, true)
	c.Ctx.WriteString("生成1000条随机数据插入到数据库中\n")
    user := new(models.Data)
    user.Id = 1213
    user.Date = time.Now()
	for i := 1; i <= 1000; i++ {
		user.Date=user.Date.Add(2*time.Hour)
		user.Value = r.Intn(10)
		o.Insert(user)
	}
}

func (c *DevController) GTdata() {
	o := orm.NewOrm()
	//生成随机数据
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	o.Using("default")
	orm.RunSyncdb("default", false, true)
    user := new(models.Data)
    user.Id = 1
    user.Date = time.Now().UTC()
	user.Value = r.Intn(10)
	o.Insert(user)
	c.Ctx.WriteString("生成1条随机数据插入到数据库中\n")

}


func (c *DevController) GRtimedata() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	datamap:=make(map[string]int64)
	datamap["x"]=time.Now().Unix()
	datamap["y"]=int64(r.Intn(10))
	c.Data["json"]=&datamap
    c.ServeJSON()
}


func (c *DevController) Get() {
	c.Ctx.WriteString("Dev")
}
