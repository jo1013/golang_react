package main

import (
	_ "bulletin_board/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	// CORS 필터 추가
	beego.InsertFilter("*", beego.BeforeExec, func(ctx *context.Context) {
		ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.ResponseWriter.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
		ctx.ResponseWriter.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	}, false)
}

func main() {
	beego.Run()
}
