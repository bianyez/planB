package controllers

import (
	"planB/models"
	"planB/util"
	"strings"
)

// 继承baseController
type AdminController struct {
	baseController
}

// Login()：用户登录
func (c *AdminController) Login() {
	if c.Ctx.Request.Method == "POST" {
		username := c.GetString("username")
		password := c.GetString("password")

		user := models.User{Username:username}
		c.o.Read(&user, "username")

		if user.Password == "" {
			c.History("密码为空", "")
		}

		if util.Md5(password) != strings.Trim(user.Password, "") {
			c.History("密码错误", "")
		}

		user.LastIp = c.getClientIp()
		user.LoginCount = user.LoginCount + 1

		if _, err := c.o.Update(&user); err != nil {
			c.History("未知登录错误", "")
		} else {
			c.History("登录成功", "/admin/main.html")
		}
		c.SetSession("user", user)
	}
	c.TplName = c.controllerName+"/login.html"
}


// 主页
func (c *AdminController) Main() {
	// c.TplName相当于http.Handle(http.FIleServer())，是用来寻找html的
	c.TplName = c.controllerName + "/main.tpl"
}
