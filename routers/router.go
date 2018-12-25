package routers

import (
	"github.com/astaxie/beego"
	"github.com/cnlh/easyProxy/controllers"
)

func init() {
	beego.Router("/", &controllers.IndexController{}, "*:Index")
	beego.AutoRouter(&controllers.IndexController{})
	beego.AutoRouter(&controllers.LoginController{})
}
