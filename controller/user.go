package controller

import (
	"fg-admin/model"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type UserController struct {
	BaseController
}

// 登录页面渲染
func (c *UserController) GetLogin() mvc.Result {
	msg := c.Session.GetFlashString("message")

	return mvc.View{
		Name: "login/login.html",
		Data: iris.Map{"message": msg},
	}
}

// 登录请求
func (c *UserController) PostLogin() {
	uname := c.Ctx.PostValueTrim("username")
	pwd := c.Ctx.PostValueTrim("password")

	token, ok, msg := models.CheckPwd(uname, pwd)
	if ok {
		c.Ctx.SetCookieKV("token", token)
		c.Ctx.Redirect("/home/index")
		return
	}

	c.Session.SetFlash("message", msg)
	c.Ctx.Redirect("/user/login")
}

// 登出
func (c *UserController) GetLogout() {
	c.Ctx.RemoveCookie("token")
	c.Ctx.Redirect("login")
}

// 用户列表页面渲染
func (c *UserController) GetList() mvc.Result {
	if !c.Ctx.IsAjax() {
		return mvc.View{
			Layout: "",
			Name:   "user/list.html",
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
	users := models.GetAllUsers(realName, "", page, limit)

	return c.pageData(users, models.GetUsersCount())

}

// 新增或修改页面渲染
func (c *UserController) GetUpsert() mvc.Result {
	uid, _ := c.Ctx.URLParamInt("id")
	var userInfo models.User
	if uid > 0 {
		userInfo = models.GetUserById(uint(uid))
	}

	roles := models.GetAllRoles("", "", 1, 10)
	return mvc.View{
		Layout: "",
		Name:   "user/upsert.html",
		Data:   iris.Map{"userInfo": userInfo, "roles": roles},
	}
}

// 新增或修改
func (c *UserController) PostUpsert() mvc.Result {
	user := new(models.User)
	id, _ := c.Ctx.PostValueInt("id")
	user.Name = c.Ctx.PostValueTrim("name")
	user.Phone = c.Ctx.PostValueTrim("phone")
	user.Email = c.Ctx.PostValueTrim("email")
	role_id, _ := c.Ctx.PostValueInt("role_id")
	user.RoleID = uint(role_id)

	if id > 0 {
		_ = models.UpdateUser(user, uint(id))
	} else {
		_ = models.CreateUser(user)
	}

	return OpSuccess
}

// 操作删除
func (c UserController) PostDelete() mvc.Result {
	id, _ := c.Ctx.PostValueInt("id")

	models.DeleteUserById(uint(id))
	return OpSuccess
}
