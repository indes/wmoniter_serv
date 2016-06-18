package routers

import (
	"wmoniter_serv/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.RtimeController{})
	beego.Router("/rtime", &controllers.RtimeController{})
	beego.Router("/historydata", &controllers.HistorydataController{})
	beego.Router("/data", &controllers.DataController{})
	beego.Router("/data/all",&controllers.DatabaseController{},"*:GetAll")
	beego.Router("/data/now",&controllers.DatabaseController{},"*:GetNow")
	beego.Router("/database",&controllers.DatabaseController{})

	//调试用路由
	beego.Router("/dev/gendata", &controllers.DevController{}, "*:GData")
	beego.Router("/dev/grtdata", &controllers.DevController{}, "*:GRtimedata")

}
