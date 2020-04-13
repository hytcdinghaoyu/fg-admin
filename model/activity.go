package models

import (
	"encoding/json"
	"fg-admin/config"
	"fmt"
)

const (
	ACTIVITY_ID_COMMON = 0
	ACTIVITY_ID_ASSIST = 1
	ACTIVITY_ID_SCORE  = 201
)

var ActivityMap = map[int32]string{ACTIVITY_ID_ASSIST: "英雄助战", ACTIVITY_ID_SCORE: "积分兑换-召唤", ACTIVITY_ID_COMMON: "通用配置"}

type ActivityCsvConfig struct {
	Name    string `json:"name" bson:"name"`       // 表名称
	Content string `json:"content" bson:"content"` // 内容
}

type ActivityInfo struct {
	Id        int32               `json:"id" bson:"id"`         // 自增ID
	Ids       []int32             `json:"ids" bson:"-"`         // 更新使用
	Sid       int32               `json:"sid" bson:"sid"`       // 服务器sid
	Sids      []int32             `json:"sids" bson:"-"`        // 增加使用
	Type      int32               `json:"type" bson:"type"`     // 活动类型
	BeginTime string              `json:"bt" bson:"bt"`         // 开始时间
	EndTime   string              `json:"et" bson:"et"`         // 结束时间
	Config    []ActivityCsvConfig `json:"config" bson:"config"` // 相关配置表
	bTime     int64                                             // 开始时间
	eTime     int64                                             // 结束时间
}

type ActivityQueryCondition struct {
	Sids      []int32 `json:"sids" bson:"sids"`   // 服务器列表
	Type      []int32 `json:"type" bson:"type"`   // 活动类型
	State     int32   `json:"state" bson:"state"` // 开启状态 0所有 1未开启 2开启 3过期
	BeginTime string  `json:"bt" bson:"bt"`       // 开始时间
	EndTime   string  `json:"et" bson:"et"`       // 结束时间
	bTime     int64                               // 开始时间
	eTime     int64                               // 结束时间
}

type ActivityListJsonRet struct {
	Code int32
	Msg  string
	Data []ActivityInfo
}

func GetAllActivity(c ActivityQueryCondition) (list []ActivityInfo) {
	ret := ActivityListJsonRet{}
	PostGmStruct("/activity/getlist", c, &ret)

	if ret.Code == STATUS_OP_FAILED {
		fmt.Println("code 0", ret.Msg)
		return nil
	}

	return ret.Data
}

func CreateActivity(a ActivityInfo) (status float64, msg string) {

	if len(a.Sids) == 0 {
		return STATUS_OP_FAILED, "参数错误"
	}

	body, err := json.Marshal(a)
	if err != nil {
		return STATUS_OP_FAILED, MSG_OP_FIALED
	}

	ret := HttpPost(config.GmServerAddr+"/activity/add", body)
	if val, ok := ret["code"]; ok {
		return val.(float64), ret["msg"].(string)
	}

	return STATUS_OP_FAILED, MSG_OP_FIALED
}

func UpdateActivity(a ActivityInfo) (status float64, msg string) {

	if len(a.Ids) == 0 {
		return STATUS_OP_FAILED, "参数错误"
	}

	body, err := json.Marshal(a)
	if err != nil {
		return STATUS_OP_FAILED, MSG_OP_FIALED
	}

	ret := HttpPost(config.GmServerAddr+"/activity/update", body)
	if val, ok := ret["code"]; ok {
		return val.(float64), ret["msg"].(string)
	}

	return STATUS_OP_FAILED, MSG_OP_FIALED
}

func DeleteActivity(ids []int32) (status float64, msg string) {
	req := make(map[string]interface{})
	req["ids"] = ids

	body, err := json.Marshal(req)
	if err != nil {
		return STATUS_OP_FAILED, MSG_OP_FIALED
	}

	ret := HttpPost(config.GmServerAddr+"/activity/delete", body)
	if val, ok := ret["code"]; ok {
		return val.(float64), ret["msg"].(string)
	}

	return STATUS_OP_FAILED, MSG_OP_FIALED
}

func LoadActivity(ids []int32) (status float64, msg string) {
	req := make(map[string]interface{})
	req["ids"] = ids

	body, err := json.Marshal(req)
	if err != nil {
		return STATUS_OP_FAILED, MSG_OP_FIALED
	}

	ret := HttpPost(config.GmServerAddr+"/activity/load", body)
	if val, ok := ret["code"]; ok {
		return val.(float64), ret["msg"].(string)
	}

	return STATUS_OP_FAILED, MSG_OP_FIALED
}
