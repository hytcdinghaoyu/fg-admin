package controller

import (
	"fg-admin/model"
	"fmt"
	"github.com/kataras/iris/v12/mvc"
	"strings"
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

	seasons := models.GetAllSeasons("", "", 1, 1000)
	list := make([]map[string]interface{}, len(seasons))
	for k, v := range seasons {
		val := make(map[string]interface{})
		val["Id"] = v.ID
		val["ActivityName"] = models.ActivityMap[v.ActivityId]
		// val["StartDate"] = time.Unix(int64(v.StartAt), 0).Format("2006-01-02 15:04:05")
		val["StartDate"] = v.StartAt
		val["EndDate"] = v.EndAt
		if v.ServerId == 1 {
			val["Server"] = "陈飞"
		} else {
			val["Server"] = "外网测试"
		}
		list[k] = val
	}

	return c.pageData(list, len(list))

}

// 新增或修改页面渲染
func (c *SeasonController) GetUpsert() mvc.Result {
	id, _ := c.Ctx.URLParamInt("id")
	var data = make(map[string]interface{})

	season, err := models.GetSeasonById(uint(id))
	if err == nil {
		data["sid"] = float64(season.ServerId)

		data["activity_id"] = "1"
		data["date_range"] = season.StartAt + " - " + season.EndAt
	} else {
		data["sid"] = float64(0)
	}

	serverList := models.GetServerList()
	if serverList != nil {
		data["serverList"] = serverList
	}

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
	serverId, _ := c.Ctx.PostValueInt("sid")
	activity_id ,_ := c.Ctx.PostValueInt("activity_id")
	dateRangeArr := strings.Split(dateRange, " - ")
	fmt.Println(dateRangeArr)
	if len(dateRangeArr) != 2 {
		return nil
	}

	cmd := make(map[string]interface{})
	cmd["act"] = "activetime"
	cmd["sid"] = serverId
	cmd["id"] = activity_id
	cmd["begin"] = dateRangeArr[0]
	cmd["end"] = dateRangeArr[1]
	status, msg := models.GmCommand(cmd)

	if status == 0 {
		season := models.Season{
			ServerId:   serverId,
			ActivityId: activity_id,
			StartAt:    dateRangeArr[0],
			EndAt:      dateRangeArr[1],
		}
		models.UpsertSeason(season)
	}

	ret := make(map[string]interface{})
	ret["status"] = status
	ret["msg"] = msg
	fmt.Println(ret)
	return mvc.Response{
		Object: ret,
	}

}
