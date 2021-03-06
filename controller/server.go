package controller

import (
	"encoding/base64"
	"encoding/json"
	"fg-admin/config"
	"fg-admin/constant"
	"fg-admin/model"
	"fg-admin/utils"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
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

	list := models.GetServerList()
	return c.pageData(list, 0)
}

// 服务器列表页面渲染
func (c *ServerController) GetStatus() mvc.Result {
	if !c.Ctx.IsAjax() {
		return mvc.View{
			Layout: "",
			Name:   "server/status.html",
		}
	}

	tp := c.Ctx.URLParamTrim("type")
	list := models.GetServerStatus(tp)
	return c.pageData(list, 0)
}

func (c *ServerController) PostList() mvc.Result {
	list := models.GetServerSelect()
	return c.pageData(list, len(list))
}

// 新增或修改页面渲染
func (c *ServerController) GetUpsert() mvc.Result {
	id, _ := c.Ctx.URLParamInt("id")

	data := models.GetServerInfo(id)
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
	ch, _ := c.Ctx.PostValueInt("ch")

	serverRecord := models.ServerInfo{
		Channel:  int32(ch),
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

	models.MemCache.Delete(constant.CACHE_SERVER_LIST)
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
	models.MemCache.Delete(constant.CACHE_SERVER_LIST)
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
	status, msg := models.GmCommand("/logic/res", cmd)

	ret := make(map[string]interface{})
	ret["status"] = status
	ret["msg"] = msg
	return mvc.Response{
		Object: ret,
	}
}

// GM命令页面渲染
func (c *ServerController) GetRes() mvc.Result {

	serverList := models.GetServerList()

	var data = make(map[string]interface{})
	if serverList != nil {
		data["serverList"] = serverList
	}

	return mvc.View{
		Layout: "",
		Name:   "server/res.html",
		Data:   data,
	}
}

// GM发送命令
func (c *ServerController) PostRes() mvc.Result {
	sid, _ := c.Ctx.PostValueInt("sid")
	cs := c.Ctx.PostValueTrim("cmd")
	uid, _ := c.Ctx.PostValueInt("uid")

	cmd := make(map[string]interface{})
	cmd["sid"] = sid
	cmd["cmd"] = cs
	cmd["uid"] = uid
	status, msg := models.GmCommand("/logic/gm", cmd)

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
	sid := c.Ctx.PostValueTrim("sid")
	uid := c.Ctx.PostValueTrim("uid")
	title := c.Ctx.PostValueTrim("title")
	content := c.Ctx.PostValueTrim("content")
	sw := c.Ctx.PostValueTrim("switch")
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
	cmd["sid"] = sid
	cmd["uid"] = uid
	cmd["title"] = title
	cmd["content"] = content
	cmd["attachment"] = attachment

	var status float64
	var msg string
	smap := models.GetServerMap()
	if sw == "on" {
		sidSli := strings.Split(sid, ",")
		for _, sid := range sidSli {
			cmd["sid"] = sid
			var smsg string
			sidi, _ := strconv.Atoi(sid)
			if sname, ok := smap[int32(sidi)]; ok {
				_, smsg = models.GmCommand("/logic/globalmail", cmd)
				fmgs := fmt.Sprintf("[%s]:%s<br>", sname ,smsg)
				msg += fmgs
			}

		}

	}else{
		status, msg = models.GmCommand("/logic/mail", cmd)
	}


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
		Data:   iris.Map{"login_srv_ip": config.LoginServerAddr, "gm_srv_ip": config.GmServerAddr, "log_srv_ip": config.LogServerAddr},
	}
}

func (c *ServerController) PostConfEdit() mvc.Result {
	config.LoginServerAddr = c.Ctx.FormValue("login_srv_ip")
	config.GmServerAddr = c.Ctx.FormValue("gm_srv_ip")
	config.LogServerAddr = c.Ctx.FormValue("log_srv_ip")

	models.MemCache.Delete(constant.CACHE_SERVER_LIST)
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

	uid, _ := c.Ctx.URLParamInt("uid")
	logs := models.GetAllOpLogs(uid, "", page, limit)
	ret := make(map[string]interface{})
	ret["data"] = logs
	ret["code"] = 0
	ret["msg"] = ""
	ret["count"] = models.GetLogsCount()
	return mvc.Response{
		Object: ret,
	}
}

// 生成兑换码
func (c *ServerController) GetCodeGen() mvc.Result {

	// 下拉选择框
	var data = make(map[string]interface{})
	data["items"] = config.ItemConf
	return mvc.View{
		Layout: "",
		Name:   "server/gencode.html",
		Data:   data,
	}

}

func (c *ServerController) PostCodeGen() string {

	name := c.Ctx.PostValueTrim("name")
	date := c.Ctx.PostValueTrim("date")
	num, _ := c.Ctx.PostValueInt("num")
	select_num, _ := c.Ctx.PostValueInt("select_num")
	var attachment string = ""
	for i := 0; i < select_num; i++ {

		ctxKI := fmt.Sprintf("item_ids[%d]", i)
		ctxKN := fmt.Sprintf("item_nums[%d]", i)

		item_id, _ := c.Ctx.PostValueInt(ctxKI)
		item_num, _ := c.Ctx.PostValueInt(ctxKN)
		if item_id > 0 && item_num > 0 {
			s := fmt.Sprintf("%d;%d|", item_id, item_num)
			attachment += s
		}
	}
	attachment = strings.TrimRight(attachment, "|")

	cmd := make(map[string]interface{})
	cmd["num"] = num
	cmd["reward"] = attachment
	cmd["name"] = name
	cmd["date"] = date
	body, _ := json.Marshal(cmd)
	ret := models.HttpPost(config.GmServerAddr+"/code/gencode", body)
	if _, ok := ret["data"]; !ok {
		return ""
	}else{
		fmt.Println(ret)
	}

	fileName := fmt.Sprintf("奖励兑换码-%s.csv", time.Now().Format("2006-01-02 15"))
	c.Ctx.Header("Content-Type", "text/csv")
	c.Ctx.Header("Content-Disposition", fmt.Sprintf("attachment;filename=%s", fileName))


	if data, ok := ret["data"]; ok && data != nil {
		data := strings.ReplaceAll(data.(string), ",", "\n")
		return data
	}

	return ""

}

// 礼包兑换码列表
func (c *ServerController) GetCodeList() mvc.Result {

	if !c.Ctx.IsAjax() {
		return mvc.View{
			Layout: "",
			Name:   "server/codelist.html",
		}
	}

	cmd := make(map[string]interface{})
	cmd["act"] = "codelist"
	body, _ := json.Marshal(cmd)
	ret := models.HttpPost(config.GmServerAddr+"/code/codelist", body)

	if data, ok := ret["data"]; ok {
		return c.pageData(data, 0)
	}

	return c.pageData([]interface{}{}, 0)
}

func (c *ServerController) PostCodeDelete() mvc.Result {
	id, _ := c.Ctx.PostValueInt("id")

	status, msg := models.DeleteCode(id)
	ret := make(map[string]interface{})
	ret["status"] = status
	ret["msg"] = msg
	return mvc.Response{
		Object: ret,
	}
}

func (c *ServerController) GetCsv() mvc.Result {
	if !c.Ctx.IsAjax() {
		return mvc.View{
			Layout: "",
			Name:   "server/csv.html",
		}
	}

	cmd := make(map[string]interface{})
	body, _ := json.Marshal(cmd)
	ret := models.HttpPost(config.GmServerAddr+"/csv/getlist", body)

	if data, ok := ret["data"]; ok {
		if dataSli, ok := data.([]interface{}); ok {
			for _, v := range dataSli {
				if m, ok := v.(map[string]interface{}); ok {
					tm := time.Unix(int64(m["time"].(float64)), 0)
					m["time"] = tm.Format("2006-01-02 15:04:05")

					if stype, ok := models.MServerType[int32(m["stype"].(float64))]; ok {
						m["stype"] = stype
					}
				}
			}
		}

		return c.pageData(data, 0)
	}
	return c.pageData([]interface{}{}, 0)
}

func (c *ServerController) GetCsvUpsert() mvc.Result {
	data := map[string]interface{}{}
	data["upload_key"] = utils.RandString(8)
	return mvc.View{
		Layout: "",
		Name:   "server/upsertCsv.html",
		Data:   data,
	}
}

func (c *ServerController) PostCsvUpsert() mvc.Result {
	sids := c.Ctx.PostValueTrim("sids")
	battle_sids := c.Ctx.PostValueTrim("battle_sids")
	stype, _ := c.Ctx.PostValueInt("stype")
	upload_key := c.Ctx.PostValueTrim("upload_key")

	if sids != "" {
		sids = strings.ReplaceAll(sids, ",", "|")
	}

	if stype == models.SERVER_TYPE_BATTLE {
		sids = battle_sids
	}

	cmd := map[string]interface{}{}
	content, ok := models.MemCache.Get(upload_key + "content")
	if ok {
		cmd["content"] = content
	}

	name, ok := models.MemCache.Get(upload_key + "name")
	if ok {
		cmd["name"] = name
	}

	cmd["sids"] = sids
	cmd["stype"] = stype

	body, err := json.Marshal(cmd)
	if err != nil {
		return mvc.Response{
			Object: nil,
		}
	}

	ret := models.HttpPost(config.GmServerAddr+"/csv/upsert", body)

	return mvc.Response{
		Object: ret,
	}
}

func (c *ServerController) PostCsvUpload() mvc.Result {
	file, info, err := c.Ctx.FormFile("file")
	upload_key := c.Ctx.PostValueTrim("upload_key")
	if err != nil {
		return c.fireError(err)
	}

	rBuf, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("[read gzip data err]: ", err)
	}

	// encode
	content := base64.RawStdEncoding.EncodeToString(rBuf)
	models.MemCache.Set(upload_key+"content", content, time.Second*300)
	models.MemCache.Set(upload_key+"name", info.Filename, time.Second*300)

	return OpSuccess
}

func (c *ServerController) PostCsvLoad() mvc.Result {

	ret := models.HttpPost(config.GmServerAddr+"/csv/load", nil)
	return mvc.Response{
		Object: ret,
	}
}

func (c *ServerController) PostCsvDelete() mvc.Result {

	id, _ := c.Ctx.PostValueInt("id")

	cmd := map[string]interface{}{}
	cmd["id"] = id
	body, _ := json.Marshal(cmd)
	ret := models.HttpPost(config.GmServerAddr+"/csv/delete", body)
	return mvc.Response{
		Object: ret,
	}
}
