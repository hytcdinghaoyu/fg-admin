package controller

import (
	"fg-admin/config"
	models "fg-admin/model"
	"fmt"
	"github.com/kataras/iris/v12/mvc"
	"strconv"
	"strings"
	"time"
)

type MonitorController struct {
	BaseController
}

// 玩家数据
func (c *MonitorController) GetPlayer() mvc.Result {
	sid, _ := c.Ctx.URLParamInt("sid")
	uid, _ := c.Ctx.URLParamInt("uid")

	data := map[string]interface{}{}
	data["sid"] = float64(sid)
	if uid > 0 {
		data["uid"] = uid
	}else {
		data["uid"] = ""
	}

	p := models.GetPlayerInfo(sid, uid)
	if p != nil {
		data["data"] = p
	}

	serverList := models.GetServerList()
	if serverList != nil {
		data["serverList"] = serverList
	}
	return mvc.View{
		Name: "monitor/player.html",
		Data: data,
	}
}

// 在线数
func (c *MonitorController) GetOnline() mvc.Result {
	var d string
	sid, _ := c.Ctx.URLParamInt("sid")
	d = c.Ctx.URLParamTrim("date")
	if d == "" {
		d = time.Now().Format("2006-01-02")
	}

	data := map[string]interface{}{}
	data["sid"] = float64(sid)
	data["date"] = d

	serverList := models.GetServerList()
	if serverList != nil {
		data["serverList"] = serverList
	}

	return mvc.View{
		Name: "monitor/online.html",
		Data: data,
	}
}

func (c *MonitorController) PostOnline() mvc.Result {
	var date string
	sid, err := c.Ctx.PostValueInt("sid")
	if err != nil {
		sid = 1
	}
	date = c.Ctx.PostValueTrim("date")
	if date == "" {
		date = time.Now().Format("2006-01-02")
	}

	loc, _ := time.LoadLocation("Local")
	theTime, _ := time.ParseInLocation("2006-01-02", date, loc)

	ret := make(map[string]interface{})
	midnight := theTime.Unix()
	xData := []string{}
	yData := []int32{}
	data := models.GetOnlineCount(sid, midnight, midnight+3600*24)
	if data, ok := data.([]interface{}); ok {
		for _, v := range data {
			if m, ok := v.(map[string]interface{}); ok {

				// if int64(m["time"].(float64)) % 300 != 0 {
				// 	continue
				// }

				if t, ok := m["time"]; ok {
					xData = append(xData, time.Unix(int64(t.(float64)), 0).Format("15:04"))
				}

				if n, ok := m["num"]; ok {
					yData = append(yData, int32(n.(float64)))
				}
			}
		}
	}

	ret["xData"] = xData
	ret["yData"] = yData
	return mvc.Response{
		Object: ret,
	}
}

// 导出在线人数csv
func (c *MonitorController) PostOnlineExport() string {
	var date string
	sid, err := c.Ctx.PostValueInt("sid")
	if err != nil {
		sid = 1
	}
	date = c.Ctx.PostValueTrim("date")
	if date == "" {
		date = time.Now().Format("2006-01-02")
	}

	loc, _ := time.LoadLocation("Local")
	theTime, _ := time.ParseInLocation("2006-01-02", date, loc)

	ret := "时间,在线人数\n"
	midnight := theTime.Unix()
	data := models.GetOnlineCount(sid, midnight, midnight+3600*24)
	if data, ok := data.([]interface{}); ok {
		for _, v := range data {
			if m, ok := v.(map[string]interface{}); ok {
				line := time.Unix(int64(m["time"].(float64)), 0).Format("15:04") + "," +
					strconv.FormatInt(int64(m["num"].(float64)), 10) + "\n"
				ret += line
			}
		}
	}

	fileName := fmt.Sprintf("在线人数统计-%s.csv", time.Now().Format("2006-01-02 15"))
	c.Ctx.Header("Content-Type", "text/csv")
	c.Ctx.Header("Content-Disposition", fmt.Sprintf("attachment;filename=%s", fileName))
	return ret
}

func (c *MonitorController) GetFlow() mvc.Result {

	if !c.Ctx.IsAjax() {
		data := map[string]interface{}{}
		serverList := models.GetServerList()
		if serverList != nil {
			data["serverList"] = serverList
		}
		return mvc.View{
			Name: "monitor/flow.html",
			Data: data,
		}
	}

	req := make(map[string]interface{})
	page, _ := c.Ctx.URLParamInt("page")
	limit, _ := c.Ctx.URLParamInt("limit")
	dateRange := c.Ctx.URLParamTrim("dateRange")

	dateRangeArr := strings.Split(dateRange, " - ")
	if len(dateRangeArr) == 2 {
		req["btime"] = dateRangeArr[0]
		req["etime"] = dateRangeArr[1]
	}

	itemId, _ := c.Ctx.URLParamInt("itemId")
	if itemId > 0 {
		req["itemId"] = itemId
	}

	uid, _ := c.Ctx.URLParamInt("uid")
	if uid > 0 {
		req["uid"] = uid
	}

	sid, _ := c.Ctx.URLParamInt("sid")
	if sid > 0 {
		req["sid"] = sid
	}

	req["skip"] = page - 1
	req["limit"] = limit
	list := models.GetLogFlow(req)
	return c.pageData(list, len(list))
}

// 在线数
func (c *MonitorController) GetPprof() mvc.Result {
	sid, _ := c.Ctx.URLParamInt("sid")
	serverName := c.Ctx.URLParamTrim("serverName")

	data := map[string]interface{}{}
	data["sid"] = float64(sid)
	data["serverName"] = serverName
	data["gmUrl"] = config.GmServerAddr

	return mvc.View{
		Name: "monitor/pprof.html",
		Data: data,
	}
}