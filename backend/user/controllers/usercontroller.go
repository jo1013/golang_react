package controllers

import (
    "your_project/backend/bulletin_board/models"
    "github.com/astaxie/beego/orm"
    "time"
)

type UserController struct {
    beego.Controller
}

func (c *UserController) Signup() {
    var user models.User
    if err := c.ParseForm(&user); err != nil {
        c.Data["json"] = map[string]string{"error": "Invalid data"}
        c.ServeJSON()
        return
    }

    // 여기에 비밀번호 해싱 로직 추가 필요

    user.CreatedAt = time.Now()
    user.UpdatedAt = time.Now()

    o := orm.NewOrm()
    _, err := o.Insert(&user)
    if err != nil {
        c.Data["json"] = map[string]string{"error": "Failed to register user"}
        c.ServeJSON()
        return
    }

    c.Data["json"] = map[string]string{"success": "User registered successfully"}
    c.ServeJSON()
}
