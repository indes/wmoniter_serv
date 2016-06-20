package routers

import (
	"wmoniter_serv/controllers"
	"github.com/astaxie/beego"
)

func init() {
	//页面控制器
	beego.Router("/", &controllers.MainController{})
	beego.Router("/historydata/?:pn", &controllers.MainController{},"*:History")
	beego.Router("/alert/?:pn", &controllers.MainController{},"*:Alert")
	beego.Router("/settings", &controllers.MainController{},"get:Settings;post:Set")

	//数据控制器
	beego.Router("/data", &controllers.DataController{})
	beego.Router("/data/all",&controllers.DatabaseController{},"*:GetAll")
	beego.Router("/data/now",&controllers.DatabaseController{},"*:GetNow")
	beego.Router("/database",&controllers.DatabaseController{})

	//调试用路由
	beego.Router("/dev/gendata/?:num", &controllers.DevController{}, "*:GData")
	beego.Router("/dev/grtdata", &controllers.DevController{}, "*:GRtimedata")
	beego.Router("/dev/gtdata", &controllers.DevController{}, "*:GTdata")

}
