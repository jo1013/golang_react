package main

import (
	_ "user/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run() // Beego 서버 실행
}
