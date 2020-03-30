package models

import (
	"encoding/json"
	"fg-admin/constant"
)

const (
	STATUS_OP_FAILED = -1
	MSG_OP_FIALED    = "操作失败"
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

func GetServerList() interface{} {

	serverList, _ := MemCache.Get(constant.CACHE_SERVER_LIST)
	if serverList == nil {
		ret := PostLogin("/serverlist", []byte(`{"act": "select"}`))
		if data, ok := ret["data"]; ok {
			MemCache.Set(constant.CACHE_SERVER_LIST, data, -1)
			return data
		}
	}
	return serverList
}

func GetServerIds() []int32 {
	serverList :=  GetServerList()
	list, ok := serverList.([]interface{})
	ids := []int32{}
	if ok {
		for _, v := range list {
			if m, ok := v.(map[string]interface{}); ok {
 				ids = append(ids, int32(m["id"].(float64)))
			}
		}
	}
	return ids
}

func GetServerSelect() []map[string]interface{} {
	serverList :=  GetServerList()
	list, ok := serverList.([]interface{})
	ret := []map[string]interface{}{}
	if ok {
		for _, v := range list {
			if m, ok := v.(map[string]interface{}); ok {
				s := map[string]interface{}{}
				s["value"] = int32(m["id"].(float64))
				s["name"] = m["nm"].(string)
				s["selected"] = true
				ret = append(ret, s)
			}
		}
	}
	return ret
}

func GetServerMap() map[int32]string {
	serverList :=  GetServerList()
	list, ok := serverList.([]interface{})
	ret := map[int32]string{}
	if ok {
		for _, v := range list {
			if m, ok := v.(map[string]interface{}); ok {
				ret[int32(m["id"].(float64))] = m["nm"].(string)
			}
		}
	}
	return ret
}

func GetServerInfo(sid int) map[string]interface{} {
	serverList :=  GetServerList()
	list, ok := serverList.([]interface{})
	ret := map[string]interface{}{}
	if ok {
		for _, v := range list {
			if m, ok := v.(map[string]interface{}); ok {
				if int(m["id"].(float64)) == sid {
					return m
				}
			}
		}
	}
	return ret
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
