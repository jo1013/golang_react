package controllers

import (
	"bulletin_board/models" // models 패키지
	"strconv"

	"github.com/astaxie/beego"     // Beego 패키지
	"github.com/astaxie/beego/orm" // ORM 패키지
)

type PostController struct {
	beego.Controller
}

func (c *PostController) Create() {
	var post models.Post
	if err := c.ParseForm(&post); err != nil {
		c.Data["json"] = map[string]string{"error": "Invalid data"}
		c.ServeJSON()
		return
	}

	o := orm.NewOrm()
	_, err := o.Insert(&post)
	if err != nil {
		c.Data["json"] = map[string]string{"error": "Failed to create post"}
		c.ServeJSON()
		return
	}

	c.Data["json"] = map[string]string{"success": "Post created successfully"}
	c.ServeJSON()
}

func (c *PostController) Get() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)

	var post models.Post
	o := orm.NewOrm()
	err := o.QueryTable("post").Filter("ID", id).One(&post)
	if err != nil {
		c.Data["json"] = map[string]string{"error": "Post not found"}
		c.ServeJSON()
		return
	}

	c.Data["json"] = post
	c.ServeJSON()
}

// Update와 Delete 메소드는 상황에 맞게 구현
func (c *PostController) Like() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)

	o := orm.NewOrm()
	if _, err := o.QueryTable("post").Filter("ID", id).Update(orm.Params{
		"Likes": orm.ColValue(orm.ColAdd, 1),
	}); err != nil {
		c.Data["json"] = map[string]string{"error": "Failed to update like"}
		c.ServeJSON()
		return
	}

	c.Data["json"] = map[string]string{"success": "Like updated successfully"}
	c.ServeJSON()
}
