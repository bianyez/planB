package routers

import (
	"planB/controllers"
	"github.com/astaxie/beego"
)

func init() {
    //beego.Router("/", &controllers.MainController{})
    // beego.AutoRouter()将AdminController的所有方法自动注册为路由了
    // 访问/admin/login.html或者/admin/login都可以跳转到Login()
    beego.AutoRouter(&controllers.AdminController{})
}
