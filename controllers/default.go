package controllers

import (
	"github.com/astaxie/beego/validation"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"wmoniter_serv/models"
	"strconv"

)

type Paginator struct {
	PerPageNums int
	MaxPages    int

	nums      int64
	pageRange []int
	pageNums  int
	page      int
}

func init() {

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
	o := orm.NewOrm()
	nowconf:=new(models.Conf)
	o.Raw("SELECT * FROM conf WHERE id = 1").QueryRow(&nowconf)
	c.Data["nowuplimit"]=nowconf.Uplimit
	c.Data["nowlowlimit"]=nowconf.Lowlimit
	c.Data["IsHome"]=1
	c.Data["pagetitle"]="实时数据"
	c.TplName = "index.tpl"
}

func (c *MainController) Rtime() {
	c.Data["IsHome"]=1
	c.Data["pagetitle"]="实时数据"
	c.TplName = "index.tpl"
}

func (c *MainController) History() {

	nowpn,_:=strconv.Atoi(c.Ctx.Input.Param(":pn"))
	o := orm.NewOrm()
	nowconf:=new(models.Conf)
	o.Raw("SELECT * FROM conf WHERE id = 1").QueryRow(&nowconf)

	qs := o.QueryTable("data")
	var maps []*models.Data
	qs.All(&maps)
	maplens:=0
    for i,_:=range maps{
		maplens=i;
    }
	maplens=maplens+1
	qs.Limit(25, 25*(nowpn-1)).All(&maps)

	c.Data["nowuplimit"]=nowconf.Uplimit
	c.Data["nowlowlimit"]=nowconf.Lowlimit
	c.Data["data"]=maps
	maxpagenums:=0
	if maplens%25==0{
		c.Data["maxpagenums"]=maplens/25
		maxpagenums=maplens/25
	}else{
		c.Data["maxpagenums"]=(maplens/25)+1
		maxpagenums=(maplens/25)+1
	}
	fmt.Println("最大页数",c.Data["maxpagenums"])
	if nowpn==0{
		nowpn=nowpn+1
	}
	var pn1 [5]int

	if nowpn<=3{
		for i:=0;i<5;i++{
			pn1[i]=i+1
		}
	}else{
		if (nowpn+3)>maxpagenums{
			for i:=0;i<5;i++{
				pn1[i]=maxpagenums-4+i
			}
		}else{
			pn1[0]=nowpn-2
			pn1[1]=nowpn-1
			pn1[2]=nowpn
			pn1[3]=nowpn+1
			pn1[4]=nowpn+2
		}
	}

	c.Data["pn1"]=pn1
	c.Data["nowpn"]=nowpn
	c.Data["IsHistory"]=1
	c.Data["pagetitle"]="历史数据"
	c.TplName = "historydata.tpl"

}


func (c *MainController) Alert() {

	var maps1 []*models.Data
	// var maps2 []*models.Data


	o := orm.NewOrm()
	nowconf:=new(models.Conf)
	o.Raw("SELECT * FROM conf WHERE id = 1").QueryRow(&nowconf)

	qs := o.QueryTable("data")
	qs.Filter("value__gt", 8).Limit(20, 30).All(&maps1)
	// qs.Filter("value__gt", nowconf.Lowlimit).Limit(20, 30).All(&maps2)

	fmt.Println(maps1)
	for i,j:=range maps1{
		fmt.Println(i,j)
    }
	c.Data["nowuplimit"]=nowconf.Uplimit
	c.Data["nowlowlimit"]=nowconf.Lowlimit
	c.Data["IsHistory"]=1
	c.Data["pagetitle"]="报警记录"
	c.Ctx.WriteString("历史数据")
}

func (c *MainController) Settings() {
	o := orm.NewOrm()
	nowconf:=new(models.Conf)
	o.Raw("SELECT * FROM conf WHERE id = 1").QueryRow(&nowconf)
	c.Data["nowuplimit"]=nowconf.Uplimit
	c.Data["nowlowlimit"]=nowconf.Lowlimit
	c.Data["IsSettigs"]=1
	c.Data["pagetitle"]="设置"
	c.TplName = "settings.tpl"
}

func (c *MainController) Set() {
	o := orm.NewOrm()
	valid := validation.Validation{}
	valid.Numeric(c.Input().Get("lowlimit"), "lowlimit")
	valid.Numeric(c.Input().Get("uplimit"), "uplimit")
	if valid.HasErrors() {
		c.Data["messagetype"]="alert-warning"
		c.Data["message"]="请输入正确的上下限"
	}else{
		lowlimit,_:=strconv.Atoi(c.Input().Get("lowlimit"))
		uplimit,_:=strconv.Atoi(c.Input().Get("uplimit"))
		fmt.Println(lowlimit,uplimit)
		if lowlimit>=uplimit||lowlimit==0||uplimit==0{
			c.Data["messagetype"]="alert-warning"
			c.Data["message"]="修改失败，请输入正确的上下限！"
		}else{
			o.Raw("UPDATE conf SET lowlimit = ? ,uplimit =? WHERE id = 1", lowlimit, uplimit).Exec()
			c.Data["messagetype"]="alert-success"
			c.Data["message"]="修改上下限成功！"
		}
	}

	nowconf:=new(models.Conf)
	o.Raw("SELECT * FROM conf WHERE id = 1").QueryRow(&nowconf)
	c.Data["nowuplimit"]=nowconf.Uplimit
	c.Data["nowlowlimit"]=nowconf.Lowlimit
	c.Data["IsSettigs"]=1
	c.Data["pagetitle"]="设置"
	c.TplName = "settings.tpl"
}

func (c *DataController) Get() {
	c.TplName = "data.tpl"
}
