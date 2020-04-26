package controller

import (
	"fg-admin/config"
	models "fg-admin/model"
	"fmt"
	"github.com/kataras/iris/v12/mvc"
	"math"
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
	} else {
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

// 在线数
func (c *MonitorController) GetAcu() mvc.Result {
	var d string
	sid, _ := c.Ctx.URLParamInt("sid")
	d = c.Ctx.URLParamTrim("date")

	data := map[string]interface{}{}
	data["sid"] = float64(sid)
	data["date"] = d

	serverList := models.GetServerList()
	if serverList != nil {
		data["serverList"] = serverList
	}

	return mvc.View{
		Name: "monitor/acu.html",
		Data: data,
	}
}

func (c *MonitorController) PostAcu() mvc.Result {
	var date string
	sid, err := c.Ctx.PostValueInt("sid")
	if err != nil {
		sid = 1
	}
	date = c.Ctx.PostValueTrim("date")
	dateArr := strings.Split(date, " - ")
	var bDate, eDate string
	if len(dateArr) == 2 {
		bDate = dateArr[0]
		eDate = dateArr[1]
	}

	loc, _ := time.LoadLocation("Local")
	bTime, _ := time.ParseInLocation("2006-01-02", bDate, loc)
	eTime, _ := time.ParseInLocation("2006-01-02", eDate, loc)

	ret := make(map[string]interface{})
	xData := []string{}
	yDataAcu := []int32{}
	yDataPcu := []int32{}
	yDataAcuPcu := []string{}
	data := models.GetAcuPcu(sid, bTime.Unix(), eTime.Unix())
	if data, ok := data.([]interface{}); ok {
		for _, v := range data {
			if m, ok := v.(map[string]interface{}); ok {

				if t, ok := m["time"]; ok {
					xData = append(xData, time.Unix(int64(t.(float64)), 0).Format("01-02"))
				}

				n, ok := m["acu"]
				if ok {
					yDataAcu = append(yDataAcu, int32(n.(float64)))
				}

				p, ok := m["pcu"]
				if ok {
					yDataPcu = append(yDataPcu, int32(p.(float64)))
				}

				ap := fmt.Sprintf("%.2f", n.(float64)/p.(float64))
				yDataAcuPcu = append(yDataAcuPcu, ap)

			}
		}
	}

	ret["xData"] = xData
	ret["yDataAcu"] = yDataAcu
	ret["yDataPcu"] = yDataPcu
	ret["yDataAcuPcu"] = yDataAcuPcu
	return mvc.Response{
		Object: ret,
	}
}

// 导出在线人数csv
func (c *MonitorController) PostAcuExport() string {
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

// 平均在线时长
func (c *MonitorController) GetOnlineavg() mvc.Result {
	var d string
	sid, _ := c.Ctx.URLParamInt("sid")
	d = c.Ctx.URLParamTrim("date")

	data := map[string]interface{}{}
	data["sid"] = float64(sid)
	data["date"] = d

	serverList := models.GetServerList()
	if serverList != nil {
		data["serverList"] = serverList
	}

	return mvc.View{
		Name: "monitor/avg.html",
		Data: data,
	}
}

func (c *MonitorController) PostOnlineavg() mvc.Result {
	var date string
	sid, err := c.Ctx.PostValueInt("sid")
	if err != nil {
		sid = 1
	}
	date = c.Ctx.PostValueTrim("date")
	dateArr := strings.Split(date, " - ")
	var bDate, eDate string
	if len(dateArr) == 2 {
		bDate = dateArr[0]
		eDate = dateArr[1]
	}

	loc, _ := time.LoadLocation("Local")
	bTime, _ := time.ParseInLocation("2006-01-02", bDate, loc)
	eTime, _ := time.ParseInLocation("2006-01-02", eDate, loc)

	ret := make(map[string]interface{})
	xData := []string{}
	yDataNew := []int32{}
	yDataActive := []int32{}
	dataNew := models.GetNewAvg(sid, bTime.Unix(), eTime.Unix())
	if data, ok := dataNew.([]interface{}); ok {
		for _, v := range data {
			if m, ok := v.(map[string]interface{}); ok {

				if t, ok := m["time"]; ok {
					xData = append(xData, time.Unix(int64(t.(float64)), 0).Format("01-02"))
				}

				newAvg, ok := m["avg"]
				if ok {
					yDataNew = append(yDataNew, int32(newAvg.(float64)))
				}

			}
		}
	}

	dataActive := models.GetActiveAvg(sid, bTime.Unix(), eTime.Unix())
	if data, ok := dataActive.([]interface{}); ok {
		for _, v := range data {
			if m, ok := v.(map[string]interface{}); ok {
				newAvg, ok := m["avg"]
				if ok {
					yDataActive = append(yDataActive, int32(newAvg.(float64)))
				}
			}
		}
	}

	ret["xData"] = xData
	ret["yDataNew"] = yDataNew
	ret["yDataActive"] = yDataActive
	return mvc.Response{
		Object: ret,
	}
}

// 在线时长分布
func (c *MonitorController) GetOnlinedis() mvc.Result {
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
		Name: "monitor/dis.html",
		Data: data,
	}
}

func (c *MonitorController) PostOnlinedis() mvc.Result {
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
	today, _ := time.ParseInLocation("2006-01-02", date, loc)

	ret := make(map[string]interface{})
	midnight := today.Unix()
	yesterday := today.AddDate(0, 0, -1).Unix()
	xData := []string{}
	yDataNew := []int32{}
	yDataActive := []int32{}
	dataNew := models.GetNewDis(sid, yesterday, midnight)
	if data, ok := dataNew.([]interface{}); ok {
		for _, v := range data {
			if m, ok := v.(map[string]interface{}); ok {

				if t, ok := m["online"]; ok {
					st := math.Floor(t.(float64) / 300)
					xData = append(xData, fmt.Sprintf("%.0f-%.0fm", st*5, st*5+5))
				}

				newAvg, ok := m["count"]
				if ok {
					yDataNew = append(yDataNew, int32(newAvg.(float64)))
				}

			}
		}
	}

	dataActive := models.GetActiveDis(sid, yesterday, midnight)
	if data, ok := dataActive.([]interface{}); ok {
		for _, v := range data {
			if m, ok := v.(map[string]interface{}); ok {
				newAvg, ok := m["count"]
				if ok {
					yDataActive = append(yDataActive, int32(newAvg.(float64)))
				}
			}
		}
	}

	ret["xData"] = xData
	ret["yDataNew"] = yDataNew
	ret["yDataActive"] = yDataActive
	return mvc.Response{
		Object: ret,
	}
}

// 在线时长分布
func (c *MonitorController) GetClientlog() mvc.Result {

	startDate := time.Now().AddDate(0, 0, -7).Format("2006-01-02")
	endDate := time.Now().Format("2006-01-02")
	if !c.Ctx.IsAjax() {
		d := c.Ctx.URLParamTrim("date")
		if d == "" {
			d = fmt.Sprintf("%s - %s", startDate, endDate)
		}

		data := map[string]interface{}{}
		data["date"] = d

		return mvc.View{
			Name: "monitor/clog.html",
			Data: data,
		}
	}

	date := c.Ctx.URLParamTrim("date")
	dateArr := strings.Split(date, " - ")
	var bDate, eDate string
	if len(dateArr) == 2 {
		bDate = dateArr[0]
		eDate = dateArr[1]
	}else{
		bDate = startDate
		eDate = endDate
	}

	loc, _ := time.LoadLocation("Local")
	bTime, _ := time.ParseInLocation("2006-01-02", bDate, loc)
	eTime, _ := time.ParseInLocation("2006-01-02", eDate, loc)
	fmt.Println(date, bTime, eTime)
	data := models.GetClientLog(bTime.Unix(), eTime.Unix())

	return c.pageData(data, 0)

}
