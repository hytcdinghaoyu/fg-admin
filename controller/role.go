package controller

import (
	"fg-admin/model"
	"github.com/kataras/iris/v12/mvc"
	"strconv"
	"strings"
)

type RoleController struct {
	BaseController
}

// 角色列表
func (c *RoleController) GetList() mvc.Result {
	if !c.Ctx.IsAjax() {
		return mvc.View{
			Layout: "",
			Name:   "role/list.html",
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

func (c *RoleController) GetUpsert() mvc.Result {

	id, _ := c.Ctx.URLParamInt("id")
	role := models.GetRoleById(uint(id))
	row := make(map[string]interface{})
	row["id"] = role.ID
	row["role_name"] = role.Name
	row["detail"] = role.Description

	ret := make(map[string]interface{})
	ret["role"] = row

	// 获取选择的树节点
	roleAuth := models.GetPermsByRoleId(uint(id))
	authId := make([]int, 0)
	for _, v := range roleAuth {
		authId = append(authId, int(v.ID))
	}
	ret["auth"] = authId

	return mvc.View{
		Layout: "",
		Name:   "role/upsert.html",
		Data:   ret,
	}
}

func (c *RoleController) PostUpsert() mvc.Result {
	role := new(models.Role)
	role.Name = c.Ctx.PostValueTrim("role_name")
	role.Description = c.Ctx.PostValueTrim("detail")
	authStr := c.Ctx.PostValueTrim("nodes_data")
	role_id, _ := c.Ctx.PostValueInt("id")

	authSli := strings.Split(authStr, ",")
	var authSliInt []uint
	for _, v := range authSli {
		aid, _ := strconv.Atoi(v)
		authSliInt = append(authSliInt, uint(aid))
	}

	if role_id <= 1000000 {
		models.UpdateRole(role, uint(role_id), authSliInt)
	}else{
		models.CreateRole(role, authSliInt)
	}

	return OpSuccess
}

func (c *RoleController) PostDelete() mvc.Result {
	id, _ := c.Ctx.PostValueInt("id")

    models.DeleteRoleById(uint(id))
	return OpSuccess
}
