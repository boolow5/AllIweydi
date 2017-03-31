package main

import (
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
	"github.com/boolow5/AllIweydi/controllers"
	"github.com/boolow5/AllIweydi/g"
	_ "github.com/boolow5/AllIweydi/routers"
)

func main() {
	g.InitEnv()
	controllers.InitLocales()
	beego.AddFuncMap("i18n", i18n.Tr)
	beego.AddFuncMap("neq", func(val1, val2 interface{}) bool {
		return val1 != val2
	})
	beego.Run()
}
