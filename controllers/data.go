package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "fmt"
	_ "math/rand"
	"wmoniter_serv/models"
    _ "time"
)

func init() {
    orm.Debug = true
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:root@(localhost:3306)/wmoniter?charset=utf8")
}

type DatabaseController struct {
	beego.Controller
}

func (c *DatabaseController) GetNow(){
	o := orm.NewOrm()
	nowdata:=new(models.Data)
	o.Raw("SELECT * FROM data ORDER BY id DESC LIMIT 1").QueryRow(&nowdata)
	datamap:=make(map[string]int64)
	datamap["x"]=nowdata.Date.Unix()+28800
	datamap["y"]=int64(nowdata.Value)
	c.Data["json"]=&datamap
    c.ServeJSON()
}

func (c *DatabaseController) GetAll() {
	o := orm.NewOrm()
	var maps []*models.Data
	qs := o.QueryTable("data")
	qs.All(&maps)
	alldata:=[][]int64{}
    for _,j:=range maps{
		alldata=append(alldata,[]int64{j.Date.Unix()*1000,int64(j.Value)})
    }
	c.Data["json"]=&alldata
    c.ServeJSON()
}

func (c *DatabaseController) Get() {
    o := orm.NewOrm()
	var maps []*models.Data
	qs := o.QueryTable("data")
	qs.All(&maps)
	alldata:=[][]int64{}
    for _,j:=range maps{
		alldata=append(alldata,[]int64{j.Date.Unix(),int64(j.Value)})
    }
	c.Data["json"]=&alldata
    c.ServeJSON()
}
