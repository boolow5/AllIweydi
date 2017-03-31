package routers

import (
	"github.com/astaxie/beego"
	"github.com/boolow5/AllIweydi/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/translate/:word/:args", &controllers.MiscAPIController{}, "*:GetTranslation")
}
