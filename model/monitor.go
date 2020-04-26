package models

import (
	"encoding/json"
	"fg-admin/config"
	"sort"
	"time"
)

const (
	CLIENT_LOG_ENTER_LOGIN = 1
	CLIENT_LOG_SDK_LOGIN = 2
	CLIENT_LOG_SDK_RET = 3

)

var (
	ClientLogTypeMap = map[int]string{
		CLIENT_LOG_ENTER_LOGIN: "进入登录界面",
		CLIENT_LOG_SDK_LOGIN: "SDK开始登录",
		CLIENT_LOG_SDK_RET: "SDK登录成功返回token",
	}
)

// 在线数
func GetClientLog(btime int64, etime int64) interface{} {

	req := make(map[string]interface{})
	req["step"] = "1|2|3"
	req["btime"] = btime
	req["etime"] = etime

	body, err := json.Marshal(req)
	if err != nil {
		return nil
	}

	ret := HttpPost(config.LogServerAddr+"/server/clientlog", body)

	c, ok := ret["code"];
	if !ok || c != float64(0) {
		return nil
	}

	if data, ok := ret["data"]; ok && data != nil {
		if list, ok := data.([]interface{}); ok {
			for _, v := range list {
				if m, ok := v.(map[string]interface{}); ok {
					if nm, ok := ClientLogTypeMap[int(m["step"].(float64))]; ok {
						m["name"] = nm
					}
				}
			}
		}
		return data
	}

	return nil
}

// 在线数
func GetOnlineCount(sid int, btime int64, etime int64) interface{} {

	req := make(map[string]interface{})
	req["sid"] = sid
	req["btime"] = btime
	req["etime"] = etime

	body, err := json.Marshal(req)
	if err != nil {
		return nil
	}

	ret := HttpPost(config.LogServerAddr+"/server/fiveonline", body)

	c, ok := ret["code"];
	if !ok || c != float64(0) {
		return nil
	}

	if data, ok := ret["data"]; ok && data != nil {
		return data
	}

	return nil
}

// ACU/PCU
func GetAcuPcu(sid int, btime int64, etime int64) interface{} {

	req := make(map[string]interface{})
	req["sid"] = sid
	req["btime"] = btime
	req["etime"] = etime

	body, err := json.Marshal(req)
	if err != nil {
		return nil
	}

	ret := HttpPost(config.LogServerAddr+"/server/apupcu", body)

	c, ok := ret["code"];
	if !ok || c != float64(0) {
		return nil
	}

	if data, ok := ret["data"]; ok && data != nil {
		sli := ret["data"].([]interface{})
		// sort by time
		sort.Slice(sli, func(i, j int) bool {
			im := sli[i].(map[string]interface{})
			jm := sli[j].(map[string]interface{})
			return im["time"].(float64) < jm["time"].(float64)
		})
		return sli
	}

	return nil
}

func GetNewAvg(sid int, btime int64, etime int64) interface{} {

	req := make(map[string]interface{})
	req["sid"] = sid
	req["btime"] = btime
	req["etime"] = etime

	body, err := json.Marshal(req)
	if err != nil {
		return nil
	}

	ret := HttpPost(config.LogServerAddr+"/server/newavg", body)

	c, ok := ret["code"];
	if !ok || c != float64(0) {
		return nil
	}

	if data, ok := ret["data"]; ok && data != nil {
		sli := data.([]interface{})
		// sort by time
		sort.Slice(sli, func(i, j int) bool {
			im := sli[i].(map[string]interface{})
			jm := sli[j].(map[string]interface{})
			return im["time"].(float64) < jm["time"].(float64)
		})
		return sli
	}

	return nil
}

func GetActiveAvg(sid int, btime int64, etime int64) interface{} {

	req := make(map[string]interface{})
	req["sid"] = sid
	req["btime"] = btime
	req["etime"] = etime

	body, err := json.Marshal(req)
	if err != nil {
		return nil
	}

	ret := HttpPost(config.LogServerAddr+"/server/activeavg", body)

	c, ok := ret["code"];
	if !ok || c != float64(0) {
		return nil
	}

	if data, ok := ret["data"]; ok && data != nil {
		sli := ret["data"].([]interface{})
		// sort by time
		sort.Slice(sli, func(i, j int) bool {
			im := sli[i].(map[string]interface{})
			jm := sli[j].(map[string]interface{})
			return im["time"].(float64) < jm["time"].(float64)
		})
		return sli
	}

	return nil
}

func GetNewDis(sid int, btime int64, etime int64) interface{} {

	req := make(map[string]interface{})
	req["sid"] = sid
	req["btime"] = btime
	req["etime"] = etime

	body, err := json.Marshal(req)
	if err != nil {
		return nil
	}

	ret := HttpPost(config.LogServerAddr+"/server/newdis", body)

	c, ok := ret["code"];
	if !ok || c != float64(0) {
		return nil
	}

	if data, ok := ret["data"]; ok && data != nil {
		sli := data.([]interface{})
		// sort by time
		sort.Slice(sli, func(i, j int) bool {
			im := sli[i].(map[string]interface{})
			jm := sli[j].(map[string]interface{})
			return im["online"].(float64) < jm["online"].(float64)
		})
		return sli
	}

	return nil
}

func GetActiveDis(sid int, btime int64, etime int64) interface{} {

	req := make(map[string]interface{})
	req["sid"] = sid
	req["btime"] = btime
	req["etime"] = etime

	body, err := json.Marshal(req)
	if err != nil {
		return nil
	}

	ret := HttpPost(config.LogServerAddr+"/server/activedis", body)

	c, ok := ret["code"];
	if !ok || c != float64(0) {
		return nil
	}

	if data, ok := ret["data"]; ok && data != nil {
		sli := data.([]interface{})
		// sort by time
		sort.Slice(sli, func(i, j int) bool {
			im := sli[i].(map[string]interface{})
			jm := sli[j].(map[string]interface{})
			return im["online"].(float64) < jm["online"].(float64)
		})
		return sli
	}

	return nil
}

func GetPlayerInfo(sid int, uid int) interface{} {
	req := make(map[string]interface{})
	req["sid"] = sid
	req["uid"] = uid

	body, err := json.Marshal(req)
	if err != nil {
		return nil
	}

	ret := PostGmArr("/logic/playerinfo", body)
	res := map[string]interface{}{}
	if len(ret) > 0 {
		for _, v := range ret {
			vv, ok := v.(map[string]interface{})
			if !ok {
				continue
			}

			if data, ok := vv["Data"]; ok {
				data, ok := data.([]interface{})
				if !ok {
					continue
				}

				data0, ok := data[0].(map[string]interface{})
				if !ok {
					continue
				}

				if vv["Table"] == "hero" {
					res["hero"] = data0["a"]
				} else if vv["Table"] == "equip" {
					res["equip"] = data0["a"]
				} else if vv["Table"] == "pet" {
					res["pet"] = data0["a"]
				} else if vv["Table"] == "item" {
					res["item"] = data0["a"]
				} else if vv["Table"] == "player" {
					data0["uid"] = int(data0["uid"].(float64))
					data0["gold"] = int(data0["gold"].(float64))
					data0["do"] = int(data0["do"].(float64))
					res["player"] = data0
				}
			}

		}

		if herolist, ok := res["hero"].([]interface{}); ok {
			for _, v := range herolist {
				if vv, ok := v.(map[string]interface{}); ok {
					vv["name"] = config.HeroConfMap[int(vv["i"].(float64))]
				}
			}
		}

		if itemlist, ok := res["item"].([]interface{}); ok {
			for _, v := range itemlist {
				if vv, ok := v.(map[string]interface{}); ok {
					vv["name"] = config.ItemConfMap[int(vv["i"].(float64))]
					vv["c"] = int(vv["c"].(float64))
				}
			}
		}

		if equiplist, ok := res["equip"].([]interface{}); ok {
			for _, v := range equiplist {
				if vv, ok := v.(map[string]interface{}); ok {
					vv["name"] = config.EquipConfMap[int(vv["i"].(float64))]
				}
			}
		}
		if petlist, ok := res["pet"].([]interface{}); ok {
			for _, v := range petlist {
				if vv, ok := v.(map[string]interface{}); ok {
					vv["name"] = config.HeroConfMap[int(vv["i"].(float64))]
				}
			}
		}
	}

	return res
}

func GetLogFlow(condition map[string]interface{}) []map[string]interface{} {
	body, err := json.Marshal(condition)
	if err != nil {
		return nil
	}

	ret := HttpPost(config.LogServerAddr+"/gamelog/query", body)
	res := []map[string]interface{}{}

	data := ret["data"]
	dataArr, ok := data.([]interface{})
	if !ok {
		return nil
	}

	for _, v := range dataArr {

		row, ok := v.(map[string]interface{})
		if !ok {
			continue
		}

		trow := map[string]interface{}{}
		for kk, vv := range row {
			if kk == "itemId" {
				trow["itemId"] = vv
			} else if kk == "count" {
				trow["count"] = vv
			} else if kk == "before" {
				trow["before"] = vv
			} else if kk == "after" {
				trow["after"] = vv
			} else if kk == "reason" {
				trow["reason"] = vv
			} else if kk == "eid" {
				trow["eid"] = vv
			} else if kk == "time" {
				trow["time"] = time.Unix(int64(vv.(float64)), 0).Format("2006-01-02 15:04")
			}
		}
		res = append(res, trow)

	}

	return res
}
