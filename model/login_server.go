package models

import (
	"encoding/json"
)

const (
	STATUS_OP_FAILED = -1
	MSG_OP_FIALED = "操作失败"
)

type ServerInfo struct {
	Id       int32  `json:"id" bson:"id"`
	Channel  int32  `json:"ch" bson:"ch"`   // 渠道
	Name     string `json:"nm" bson:"nm"`   // 名字
	GateAddr string `json:"gip" bson:"gip"` // 网关地址
	LogAddr  string `json:"lip" bson:"lip"` // 战斗日志服ip
	Status   int32  `json:"st" bson:"st"`   // 状态：拥挤、忙碌、空闲等
	Flag     int32  `json:"fg" bson:"fg"`   // 标识是否推荐服等
	Level    int32  `json:"lv" bson:"lv"`   // 该服务器对应玩家是否可见的权限

	Action string `json:"act"`
}

// 获取服务器列表
func GetServerList() interface{} {

	serverList, _ := MemCache.Get("serverlist")
	if serverList == nil {
		ret := PostLogin("/serverlist", []byte(`{"act": "select"}`))
		if data, ok := ret["data"]; ok {
			setCache(data.([]interface{}))
			return data
		}
	}
	return serverList.(map[int]interface{})
}

// 新增或修改服务器
func UpsertServer(s ServerInfo) float64 {
	body, err := json.Marshal(s)
	if err != nil {
		return STATUS_OP_FAILED
	}

	ret := PostLogin("/serverlist", body)
	if val, ok := ret["code"]; ok {
		return val.(float64)
	}
	return STATUS_OP_FAILED
}

func setCache(m []interface{}) {
	c := make(map[int]interface{})
	for _, v := range m {
		k := v.(map[string]interface{})["id"].(float64)
		c[int(k)] = v
	}
	MemCache.Set("serverlist", c, -1)
}

// 删除服务器
func DeleteServer(id int) (status float64, msg string) {

	req := make(map[string]interface{})
	req["id"] = id
	req["act"] = "delete"

	body, err := json.Marshal(req)
	if err != nil {
		return STATUS_OP_FAILED, MSG_OP_FIALED
	}

	ret := PostLogin("/serverlist", body)
	if val, ok := ret["code"]; ok {
		return val.(float64), ret["msg"].(string)
	}

	return STATUS_OP_FAILED, MSG_OP_FIALED
}

// gm命令
func GmCommand(cmd map[string]interface{}) (status float64, msg string) {

	body, err := json.Marshal(cmd)
	if err != nil {
		return STATUS_OP_FAILED, MSG_OP_FIALED
	}

	ret := PostGm("", body)
	if code, ok := ret["code"]; ok {
		return code.(float64), ret["msg"].(string)
	}

	return STATUS_OP_FAILED, MSG_OP_FIALED
}
