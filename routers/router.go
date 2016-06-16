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

}
