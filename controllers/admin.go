package controllers

import (
	"fmt"
	"planB/models"
	"planB/util"
	"strings"
)

// 继承baseController
type AdminController struct {
	baseController
}

// 接口：用户登录
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
			fmt.Println("输入的密码： ", password)
			fmt.Println("编码后的密码： ", util.Md5(password))
			fmt.Println("Trim后的db密码： ", strings.Trim(user.Password, ""))
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


// 接口：登出
func (c *AdminController) Logout() {
	c.DestroySession()
	c.History("退出登录", "/admin/login.html")
}


// 接口：网站配置信息
func (c *AdminController) Config() {
	var result []*models.Config    // Config类型的切片

	c.o.QueryTable(new(models.Config).TableName()).All(&result)

	options := make(map[string]string)    // 字典
	mp := make(map[string]*models.Config) // string类型的key，Config类型的value的字典

	for _, v := range result {
		options[v.Name] = v.Value
		mp[v.Name] = v
	}

	if c.Ctx.Request.Method == "POST" {
		keys := []string{"url", "title", "keywords", "description", "email", "start", "qq"}
		for _, key := range keys {
			val := c.GetString(key)
			if _, ok := mp[key]; !ok {
				options[key] = val
				c.o.Insert(&models.Config{Name: key, Value: val})
			} else {
				opt := mp[key]
				if _, err := c.o.Update(&models.Config{Id: opt.Id, Name: opt.Name, Value:val}); err != nil {
					continue
				}
			}
		}
		c.History("数据设置成功", "")
	}
	c.Data["config"] = options
	c.TplName = c.controllerName + "/config.html"
}


// 主页
func (c *AdminController) Main() {
	// c.TplName相当于http.Handle(http.FIleServer())，是用来寻找html的
	c.TplName = c.controllerName + "/main.tpl"
}
