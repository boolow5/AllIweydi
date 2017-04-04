package routers

import (
	"github.com/astaxie/beego"
	"github.com/boolow5/AllIweydi/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.AuthController{}, "get:Get;post:Post")
	beego.Router("/logout", &controllers.AuthController{}, "*:Logout")
	beego.Router("/register", &controllers.AuthController{}, "get:GetRegister;post:PostRegister")
	beego.Router("/translate/:word/:args", &controllers.MiscAPIController{}, "*:GetTranslation")
}
