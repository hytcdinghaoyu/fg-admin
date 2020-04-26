package controller

import (
	"encoding/json"
	"fg-admin/config"
	models "fg-admin/model"
	"github.com/kataras/iris/v12/mvc"
)

type WordsController struct {
	BaseController
}

// 列表
func (c *WordsController) GetList() mvc.Result {

	cmd := make(map[string]interface{})
	body, _ := json.Marshal(cmd)
	res := models.HttpPost(config.GmServerAddr+"/shieldingwords/getlist", body)

	ret := map[string]interface{}{}
	wordsStr := ""
	if data, ok := res["data"]; ok {
		if wordsList, ok := data.([]interface{}); ok {
			for _, v := range wordsList {
				wordsStr += v.(string) + ","
			}
		}
	}

	ret["words"] = wordsStr
	return mvc.View{
		Layout: "",
		Name:   "server/wordsList.html",
		Data:   ret,
	}
}

// 新增/修改
func (c *WordsController) PostUpsert() mvc.Result {
	words := c.Ctx.PostValueTrim("words")

	cmd := map[string]interface{}{}
	cmd["words"] = words

	body, err := json.Marshal(cmd)
	if err != nil {
		return mvc.Response{
			Object: nil,
		}
	}

	ret := models.HttpPost(config.GmServerAddr+"/shieldingwords/add", body)

	return mvc.Response{
		Object: ret,
	}
}

// 清除
func (c *WordsController) PostClear() mvc.Result {
	ret := models.HttpPost(config.GmServerAddr+"/shieldingwords/clear", nil)
	return mvc.Response{
		Object: ret,
	}
}
