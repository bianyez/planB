package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"planB/models"
	"strings"
)


type baseController struct {
	beego.Controller
	o orm.Ormer
	controllerName string
	actionName     string
}


// Prepare()，用来验证用户是否登录
func (p *baseController) Prepare() {
	controllerName, actionName := p.GetControllerAndAction()

	p.controllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])
	p.actionName = strings.ToLower(actionName)

	p.o = orm.NewOrm()
	if strings.ToLower(p.controllerName) == "admin" && strings.ToLower(p.actionName) != "login" {
		if p.GetSession("user") == nil {
			p.History("未登录", "/admin/login")
			// p.Ctx.WriteString(p.controllerName +"==="+ p.actionName)
		}
	}

	// 初始化前台页面相关元素
	if strings.ToLower(p.controllerName) == "planB" {
		p.Data["actionName"] = strings.ToLower(actionName)
		var result []*models.Config
		p.o.QueryTable(new(models.Config).TableName()).All(&result)

		configs := make(map[string]string)
		for _, v := range result {
			configs[v.Name] = v.Value
		}
		p.Data["config"] = configs
	}
}

// 做跳转的逻辑展示
func (p *baseController) History(msg string, url string) {
	if url == "" {
		p.Ctx.WriteString("<script>alert('"+msg+"');window.history.go(-1);</script>")
		p.StopRun()
	} else {
		p.Redirect(url, 302)
	}
}

// 获取用户ip
func (p *baseController) getClientIp() string {
	s := strings.Split(p.Ctx.Request.RemoteAddr, ":")
	return s[0]
}
