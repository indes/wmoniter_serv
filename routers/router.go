package routers

import (
	"wmoniter_serv/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/rtime", &controllers.RtimeController{})
}
