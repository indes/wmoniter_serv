package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	_ "math/rand"
	"wmoniter_serv/models"
    _ "time"
)

func init() {
    orm.Debug = true
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:root@(10.64.70.45:3306)/wmoniter?charset=utf8")
}


type DatabaseController struct {
	beego.Controller
}
func (c *DatabaseController) GetNow(){
	c.Ctx.WriteString("hello")
}


func (c *DatabaseController) GetAll() {
	o := orm.NewOrm()
	var maps []*models.Data
    // var maps []orm.Params
	qs := o.QueryTable("data")
    // o.Raw("SELECT * FROM data").Values(&maps)
    // fmt.Println("&v",&maps[0])
	qs.All(&maps)
	alldata:=[][]int64{}
    // alldatamap:=make(map[string]string)
	// alldata:=[]{}
    for _,j:=range maps{
		//
        // fmt.Println(i,j)
		// fmt.Println(j.Date.Unix())
		alldata=append(alldata,[]int64{j.Date.Unix(),int64(j.Value)})

    }
	c.Data["json"]=&alldata
    c.ServeJSON()

}

func (c *DatabaseController) Get() {
    o := orm.NewOrm()

	//生成随机数据
	// r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// o.Using("default")
	// orm.RunSyncdb("default", false, true)
	// c.Ctx.WriteString("123\n")
    // user := new(models.Data)
    // user.Id = 1213
    // user.Date = time.Now()
	// for i := 1; i <= 1000; i++ {
	// 	user.Date=user.Date.Add(time.Hour)
	// 	user.Value = r.Intn(10)
	// 	o.Insert(user)
	// }

	// u :=models.Data{Id: 6}
	// qs := o.QueryTable("data")
    // qs.Filter("id", 1)
	// fmt.Printf("ERR: %v\n", qs)
	// c.Ctx.WriteString(u.Date.Format("2006-01-02 15:04:05:00"))

	var maps []*models.Data
    // var maps []orm.Params
	qs := o.QueryTable("data")
    // o.Raw("SELECT * FROM data").Values(&maps)
    // fmt.Println("&v",&maps[0])
	qs.All(&maps)
	alldata:=[][]int64{}
    // alldatamap:=make(map[string]string)
	// alldata:=[]{}
    for i,j:=range maps{

        fmt.Println(i,j)
		fmt.Println(j.Date.Unix())
		alldata=append(alldata,[]int64{j.Date.Unix(),int64(j.Value)})

    }
	c.Data["json"]=&alldata
    c.ServeJSON()
}
