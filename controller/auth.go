package controller

import (
	"fg-admin/model"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type AuthController struct {
	BaseController
}

// 权限列表页面
func (c *AuthController) GetList() mvc.Result {

	if !c.Ctx.IsAjax() {
		authList := models.GetAllPermissions("", "sort ASC", 1, 1000)

		var menuList []models.Permission
		for _, v := range authList {
			if v.Pid == 0 {
				menuList = append(menuList, *v)
			}
		}

		return mvc.View{
			Layout: "",
			Name:   "auth/list.html",
			Data:   iris.Map{"menuList": menuList},
		}
	}

	// 列表
	page, err := c.Ctx.URLParamInt("page")
	if err != nil {
		page = 1
	}
	limit, err := c.Ctx.URLParamInt("limit")
	if err != nil {
		limit = 30
	}

	realName := c.Ctx.URLParamTrim("realName")
	roles := models.GetAllRoles(realName, "", page, limit)
	return c.pageData(roles, len(roles))

}

func (c *AuthController) PostTree() mvc.Result {

	authList := models.GetAllPermissions("", "sort ASC", 1, 1000)
	list := make(map[int]map[string]interface{}, len(authList))
	for _, v := range authList {
		row := make(map[string]interface{})
		row["id"] = v.ID
		row["title"] = v.Name
		row["spread"] = true
		row["pid"] = v.Pid

		if v.Pid == 0 {
			id := int(v.ID)

			if _, ok := list[id]; !ok {
				list[id] = row
				list[id]["children"] = make([]map[string]interface{}, 0)
			}

			list[id]["id"] = row["id"]
			list[id]["title"] = row["title"]
			list[id]["spread"] = row["spread"]
			list[id]["pid"] = row["pid"]
		} else {
			if _, ok := list[v.Pid]; !ok {
				list[v.Pid] = make(map[string]interface{}, 0)
				list[v.Pid]["children"] = make([]map[string]interface{}, 0)
			}
			list[v.Pid]["children"] = append(list[v.Pid]["children"].([]map[string]interface{}), row)
		}

	}

	return c.pageData(list, len(list))

}

func (c *AuthController) PostNode() mvc.Result {
	id, _ := c.Ctx.PostValueInt("id")
	result := models.GetPermissionById(uint(id))
	// if err == nil {
	// 	self.ajaxMsg(err.Error(), MSG_ERR)
	// }
	row := make(map[string]interface{})
	row["id"] = result.ID
	row["pid"] = result.Pid
	row["auth_name"] = result.Name
	row["auth_url"] = result.AuthPath
	row["sort"] = result.IsShow
	row["is_show"] = result.IsShow
	row["icon"] = result.Icon
	row["sort"] = result.Sort

	return c.pageData(row, len(row))
}

func (c *AuthController) PostUpsert() mvc.Result {

	auth := new(models.Permission)

	id, _ := c.Ctx.PostValueInt("id")
	auth.Pid, _ = c.Ctx.PostValueInt("pid")
	auth.Name = c.Ctx.PostValueTrim("auth_name")
	auth.AuthPath = c.Ctx.PostValueTrim("auth_url")
	// auth.Sort, _ = self.GetInt("sort")
	auth.IsShow, _ = c.Ctx.PostValueInt("is_show")
	auth.Icon = c.Ctx.PostValueTrim("icon")
	auth.Sort, _ = c.Ctx.PostValueInt("sort")

	if id > 0 {
		models.UpdatePermission(auth, uint(id))
	} else {
		models.CreatePermission(auth)
	}

	return OpSuccess
}

func (c *AuthController) PostDelete() mvc.Result {
	id, _ := c.Ctx.PostValueInt("id")
	if id > 0 {
		models.DeletePermissionById(uint(id))
	}

	return OpSuccess
}
