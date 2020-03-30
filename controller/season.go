package controller

import (
	"fg-admin/model"
	"fg-admin/utils"
	"fmt"
	"github.com/kataras/iris/v12/mvc"
	"io/ioutil"
	"strings"
	"time"
)

type SeasonController struct {
	BaseController
}

// 列表页面渲染
func (c *SeasonController) GetList() mvc.Result {

	if !c.Ctx.IsAjax() {
		return mvc.View{
			Layout: "",
			Name:   "season/list.html",
		}
	}

	condition := models.ActivityQueryCondition{}
	typeIds := c.Ctx.URLParamTrim("type")
	sids := c.Ctx.URLParamTrim("sid")
	status, _ := c.Ctx.URLParamInt("status")
	dateRange := c.Ctx.URLParamTrim("dateRange")

	dateRangeArr := strings.Split(dateRange, " - ")
	if len(dateRangeArr) == 2 {
		condition.BeginTime = dateRangeArr[0]
		condition.EndTime = dateRangeArr[1]
	}

	if typeIds == "" {
		condition.Type = []int32{models.ACTIVITY_ID_ASSIST, models.ACTIVITY_ID_SCORE}
	} else {
		condition.Type = utils.StrSplit(typeIds, ",")
	}

	if sids != "" {
		condition.Sids = utils.StrSplit(sids, ",")
	} else {
		condition.Sids = models.GetServerIds()
	}

	condition.State = int32(status)
	acts := models.GetAllActivity(condition)

	list := make([]map[string]interface{}, len(acts))
	loc, _ := time.LoadLocation("Local")
	nowTime := time.Now().Unix()
	serverMap := models.GetServerMap()
	for k, v := range acts {
		val := make(map[string]interface{})
		val["Id"] = v.Id

		if _, ok := models.ActivityMap[v.Type]; ok {
			val["ActivityName"] = models.ActivityMap[v.Type]
		}

		val["ActivityType"] = v.Type
		val["StartDate"] = v.BeginTime
		val["EndDate"] = v.EndTime

		beginTime, _ := time.ParseInLocation("2006-01-02 15:04:05", v.BeginTime, loc)
		endTime, _ := time.ParseInLocation("2006-01-02 15:04:05", v.EndTime, loc)
		if nowTime < beginTime.Unix() {
			val["StatusText"] = "未开始"
		} else if nowTime > beginTime.Unix() && nowTime < endTime.Unix() {
			val["StatusText"] = "进行中"
		} else if nowTime > endTime.Unix() {
			val["StatusText"] = "已结束"
		}

		if serverName, ok := serverMap[v.Sid]; ok {
			val["Server"] = serverName
		}

		val["FileList"] = v.Config

		list[k] = val
	}

	return c.pageData(list, len(list))

}

// 新增或修改页面渲染
func (c *SeasonController) GetUpsert() mvc.Result {
	var data = make(map[string]interface{})

	data["upload_key"] = utils.RandString(8)
	data["activity_map"] = models.ActivityMap
	return mvc.View{
		Layout: "",
		Name:   "season/upsert.html",
		Data:   data,
	}
}

// 新增或修改
func (c *SeasonController) PostUpsert() mvc.Result {
	dateRange := c.Ctx.PostValueTrim("date-range")
	sids := c.Ctx.PostValueTrim("select")
	activity_type, _ := c.Ctx.PostValueInt("activity_type")
	upload_key := c.Ctx.PostValueTrim("upload_key")
	dateRangeArr := strings.Split(dateRange, " - ")

	act := models.ActivityInfo{}
	act.Sids = utils.StrSplit(sids, ",")
	act.Type = int32(activity_type)
	if len(dateRangeArr) == 2 {
		act.BeginTime = dateRangeArr[0]
		act.EndTime = dateRangeArr[1]
	}

	mem, ok := models.MemCache.Get(upload_key)
	if ok {
		if m, ok := mem.(map[string]string); ok {
			for k, v := range m {
				csv := models.ActivityCsvConfig{Name: k, Content: v}
				act.Config = append(act.Config, csv)
			}
		}
	}

	status, msg := models.CreateActivity(act)
	ret := make(map[string]interface{})
	ret["status"] = status
	ret["msg"] = msg
	fmt.Println(ret)
	return mvc.Response{
		Object: ret,
	}

}

func (c *SeasonController) PostDelete() mvc.Result {
	ids := c.Ctx.PostValueTrim("id")
	ids = strings.TrimRight(ids, ",")
	idsArr := utils.StrSplit(ids, ",")

	status, msg := models.DeleteActivity(idsArr)
	ret := make(map[string]interface{})
	ret["status"] = status
	ret["msg"] = msg
	return mvc.Response{
		Object: ret,
	}
}

func (c *SeasonController) PostUpload() mvc.Result {
	file, info, err := c.Ctx.FormFile("file")
	upload_key := c.Ctx.PostValueTrim("upload_key")
	if err != nil {
		return c.fireError(err)
	}

	b, err := ioutil.ReadAll(file)
	mem, ok := models.MemCache.Get(upload_key)
	if ok {
		if v, ok := mem.(map[string]string); ok {
			v[info.Filename] = string(b)
			models.MemCache.Set(upload_key, v, time.Second*300)
		}
	} else {
		v := map[string]string{}
		v[info.Filename] = string(b)
		models.MemCache.Set(upload_key, v, time.Second*300)
	}

	return OpSuccess
}

func (c *SeasonController) GetEdit() mvc.Result {
	Id, _ := c.Ctx.URLParamInt("id")
	serverName := c.Ctx.URLParamTrim("server_name")
	actName := c.Ctx.URLParamTrim("act_name")
	startDate := c.Ctx.URLParamTrim("start_date")
	endDate := c.Ctx.URLParamTrim("end_date")

	data := map[string]interface{}{}
	data["Id"] = Id
	data["serverName"] = serverName
	data["actName"] = actName
	data["startDate"] = startDate
	data["endDate"] = endDate
	data["upload_key"] = utils.RandString(8)
	return mvc.View{
		Layout: "",
		Name:   "season/edit.html",
		Data:   data,
	}
}

func (c *SeasonController) PostEdit() mvc.Result {
	id, _ := c.Ctx.PostValueInt("id")
	dateRange := c.Ctx.PostValueTrim("date-range")
	dateRangeArr := strings.Split(dateRange, " - ")
	upload_key := c.Ctx.PostValueTrim("upload_key")

	act := models.ActivityInfo{}
	act.Ids = []int32{int32(id)}
	if len(dateRangeArr) == 2 {
		act.BeginTime = dateRangeArr[0]
		act.EndTime = dateRangeArr[1]
	}

	mem, ok := models.MemCache.Get(upload_key)
	if ok {
		if m, ok := mem.(map[string]string); ok {
			for k, v := range m {
				csv := models.ActivityCsvConfig{Name: k, Content: v}
				act.Config = append(act.Config, csv)
			}
		}
	}

	status, msg := models.UpdateActivity(act)
	ret := make(map[string]interface{})
	ret["status"] = status
	ret["msg"] = msg
	return mvc.Response{
		Object: ret,
	}

}

// 加载配置
func (c *SeasonController) PostLoad() mvc.Result {
	ids := []int32{}
	status, msg := models.LoadActivity(ids)
	ret := make(map[string]interface{})
	ret["status"] = status
	ret["msg"] = msg
	return mvc.Response{
		Object: ret,
	}
}
