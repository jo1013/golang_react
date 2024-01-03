package routers

import (
	"user/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/user", &controllers.UserController{})
}
