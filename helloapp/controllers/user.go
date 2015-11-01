package controllers

import (
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Profile() {
	c.TplNames = "user/profile.html"
	c.Data["userid"] = "liuyi"
	c.Data["tag"] = "hello liuyi"
	c.Data["hobby"] = []string{"hhhh", "jjj"}
}

// func init() {
// 	beego.Router("/", &controllers.Maincontroller{})

// 	beego.Router("/user/profile", &controllers.Usercontroller{}, "get")
// }
