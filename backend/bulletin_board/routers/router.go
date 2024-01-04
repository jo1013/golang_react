package routers

import (
	"bulletin_board/controllers"

	"github.com/astaxie/beego"
)

func init() {

	// 게시글 관련 라우트
	beego.Router("/post/", &controllers.PostController{}, "post:Create")
	beego.Router("/post/:id", &controllers.PostController{}, "get:Get;put:Update;delete:Delete")
	beego.Router("/post/:id/like", &controllers.PostController{}, "post:Like")

}
