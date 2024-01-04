package routers

import (
	"user/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/auth/signup", &controllers.UserController{}, "post:Signup") // POST 메서드를 위한 라우팅 추가
}
