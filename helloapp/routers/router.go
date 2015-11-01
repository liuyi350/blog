package routers

import (
	"blog/helloapp/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/user/profile", &controllers.UserController{}, `get:Profile`)
}
