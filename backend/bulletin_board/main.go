package main

import (
    "github.com/astaxie/beego"
    _ "bulletin_board/routers"
)

func main() {
    // Set up any necessary Beego settings here

    beego.Run() // Start the Beego server
}
