package controller

import (
	"fg-admin/config"
	"fg-admin/model"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"io"
	"os"
	"strings"
)

type ServerController struct {
	BaseController
}

// 服务器列表页面渲染
func (c *ServerController) GetList() mvc.Result {
	if !c.Ctx.IsAjax() {
		return mvc.View{
			Layout: "",
			Name:   "server/list.html",
		}
	}

	list, ok := models.GetServerList().(map[int]interface{})
	fmt.Println(list)
	if !ok {
		list = nil
	}

	return c.pageData(list, len(list))
}

// 新增或修改页面渲染
func (c *ServerController) GetUpsert() mvc.Result {
	id, _ := c.Ctx.URLParamInt("id")
	serverList := models.GetServerList()

	var data = make(map[string]interface{})
	if serverList == nil {
		data["server"] = nil
	} else {
		if res, ok := serverList.(map[int]interface{})[id]; ok {
			data["server"] = res
		}
	}

	return mvc.View{
		Layout: "",
		Name:   "server/upsert.html",
		Data:   data,
	}
}

// 新增或修改
func (c *ServerController) PostUpsert() mvc.Result {
	serverName := c.Ctx.PostValueTrim("nm")
	GateIP := c.Ctx.PostValueTrim("gip")
	LogIP := c.Ctx.PostValueTrim("lip")
	st, _ := c.Ctx.PostValueInt("st")
	flag, _ := c.Ctx.PostValueInt("fg")
	level, _ := c.Ctx.PostValueInt("lv")
	id, _ := c.Ctx.PostValueInt("id")

	serverRecord := models.ServerInfo{
		Channel:  0,
		Name:     serverName,
		GateAddr: GateIP,
		LogAddr:  LogIP,
		Status:   int32(st),
		Flag:     int32(flag),
		Level:    int32(level),
		Action:   "upsert",
	}

	if id > 0 {
		serverRecord.Id = int32(id)
	}

	status := models.UpsertServer(serverRecord)
	ret := make(map[string]interface{})
	ret["status"] = status
	ret["msg"] = ""
	return mvc.Response{
		Object: ret,
	}

}

func (c *ServerController) PostDelete() mvc.Result {
	id, _ := c.Ctx.PostValueInt("id")

	status, msg := models.DeleteServer(id)
	ret := make(map[string]interface{})
	ret["status"] = status
	ret["msg"] = msg
	return mvc.Response{
		Object: ret,
	}
}

// GM命令页面渲染
func (c *ServerController) GetCmd() mvc.Result {

	serverList := models.GetServerList()

	var data = make(map[string]interface{})
	if serverList != nil {
		data["serverList"] = serverList
	}

	// 下拉选择框
	data["items"] = config.ItemConf

	return mvc.View{
		Layout: "",
		Name:   "server/cmd.html",
		Data:   data,
	}
}

// GM发送命令
func (c *ServerController) PostCmd() mvc.Result {
	sid, _ := c.Ctx.PostValueInt("sid")
	tp, _ := c.Ctx.PostValueInt("type")
	num, _ := c.Ctx.PostValueInt("num")
	uid, _ := c.Ctx.PostValueInt("uid")

	cmd := make(map[string]interface{})
	cmd["act"] = "res"
	cmd["sid"] = sid
	cmd["type"] = tp
	cmd["num"] = num
	cmd["uid"] = uid
	status, msg := models.GmCommand(cmd)

	ret := make(map[string]interface{})
	ret["status"] = status
	ret["msg"] = msg
	return mvc.Response{
		Object: ret,
	}
}

// GM发送邮件页面渲染
func (c *ServerController) GetMail() mvc.Result {
	sid, err := c.Ctx.URLParamInt("id")
	var data = make(map[string]interface{})
	if err == nil {
		data["sid"] = float64(sid)
	} else {
		data["sid"] = float64(0)
	}
	serverList := models.GetServerList()
	if serverList != nil {
		data["serverList"] = serverList
	}

	// 下拉选择框
	data["items"] = config.ItemConf

	return mvc.View{
		Layout: "",
		Name:   "server/mail.html",
		Data:   data,
	}
}

// GM发送邮件
func (c *ServerController) PostMail() mvc.Result {
	sid, _ := c.Ctx.PostValueInt("sid")
	uid := c.Ctx.PostValueTrim("uid")
	title := c.Ctx.PostValueTrim("title")
	content := c.Ctx.PostValueTrim("content")
	select_num, _ := c.Ctx.PostValueInt("select_num")

	var attachment string = ""
	for i := 0; i < select_num; i++ {

		ctxKI := fmt.Sprintf("item_ids[%d]", i)
		ctxKN := fmt.Sprintf("item_nums[%d]", i)

		item_id, _ := c.Ctx.PostValueInt(ctxKI)
		item_num, _ := c.Ctx.PostValueInt(ctxKN)
		s := fmt.Sprintf("%d;%d|", item_id, item_num)
		attachment += s
	}
	attachment = strings.TrimRight(attachment, "|")

	cmd := make(map[string]interface{})
	cmd["act"] = "mail"
	cmd["sid"] = sid
	cmd["uid"] = uid
	cmd["title"] = title
	cmd["content"] = content
	cmd["attachment"] = attachment
	status, msg := models.GmCommand(cmd)

	ret := make(map[string]interface{})
	ret["status"] = status
	ret["msg"] = msg
	return mvc.Response{
		Object: ret,
	}
}

func (c *ServerController) GetConfEdit() mvc.Result {
	return mvc.View{
		Layout: "",
		Name:   "server/confEdit.html",
		Data:   iris.Map{"login_srv_ip": config.LoginServerAddr, "gm_srv_ip": config.GmServerAddr},
	}
}

func (c *ServerController) PostConfEdit() mvc.Result {
	config.LoginServerAddr = c.Ctx.FormValue("login_srv_ip")
	config.GmServerAddr = c.Ctx.FormValue("gm_srv_ip")

	return OpSuccess
}

func (c *ServerController) GetUpload() mvc.Result {
	return mvc.View{
		Layout: "",
		Name:   "server/upload.html",
	}
}

func (c *ServerController) PostUpload() mvc.Result {

	file, info, err := c.Ctx.FormFile("file")
	if err != nil {
		return c.fireError(err)
	}

	defer file.Close()
	fname := info.Filename
	out, err := os.OpenFile("./upload/"+fname,
		os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return c.fireError(err)
	}
	defer out.Close()
	if _, err := io.Copy(out, file); err == nil {
		config.ItemConf = config.NewItemConf()
	}

	return OpSuccess

}

// 用户列表页面渲染
func (c *ServerController) GetOperateList() mvc.Result {
	if !c.Ctx.IsAjax() {
		return mvc.View{
			Layout: "",
			Name:   "server/operateList.html",
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
	logs := models.GetAllOpLogs(realName, "", page, limit)
	ret := make(map[string]interface{})
	ret["data"] = logs
	ret["code"] = 0
	ret["msg"] = ""
	ret["count"] = models.GetLogsCount()
	return mvc.Response{
		Object: ret,
	}
}