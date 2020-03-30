package controller

import (
	"fg-admin/constant"
	"fg-admin/model"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type HomeController struct {
	BaseController
}

func (c *HomeController) AfterActivation(a mvc.AfterActivation) {
	if a.Singleton() {
		panic("basicController should be stateless,a request-scoped,we have a 'Session' which depends on the context.")
	}
}

// 主页侧边栏菜单
func (c *HomeController) GetIndex() mvc.Result {

	username := c.Ctx.Values().GetString(constant.CTX_USERNAME)
	roleId, _ := c.Ctx.Values().GetUint(constant.CTX_ROLEID)
	pemList := models.GetPermsByRoleId(roleId)

	var mainMenu, subMenu []map[string]interface{}
	for _, v := range pemList {
		if v.IsShow == 0 {
			continue
		}
		row := make(map[string]interface{})
		row["Id"] = v.ID
		// row["Sort"] = v.Sort
		row["AuthName"] = v.Name
		row["AuthUrl"] = v.AuthPath
		row["Icon"] = v.Icon
		row["Pid"] = v.Pid
		if v.Pid == 0 {
			mainMenu = append(mainMenu, row)
		} else {
			subMenu = append(subMenu, row)
		}
	}

	return mvc.View{
		Name: "public/menu.html",
		Data: iris.Map{"MainMenu": mainMenu, "SubMenu": subMenu, "SiteName": "FG-ADMIN", "Username": username},
	}
}

// 默认显示的标签页
func (c *HomeController) GetMenu() mvc.Result {
	return mvc.View{
		Name:   "home/index.html",
		Layout: "",
	}
}
