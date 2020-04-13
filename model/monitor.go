package models

import (
	"encoding/json"
	"fg-admin/config"
	"time"
)

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

	ret := HttpPost(config.GmServerAddr+"/server/online", body)
	if _, ok := ret["data"]; ok {
		return ret["data"]
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
				}else if vv["Table"] == "equip" {
					res["equip"] = data0["a"]
				}else if vv["Table"] == "pet" {
					res["pet"] = data0["a"]
				}else if vv["Table"] == "item" {
					res["item"] = data0["a"]
				}else if vv["Table"] == "player" {
					data0["uid"] = int(data0["uid"].(float64))
					data0["gold"] = int(data0["gold"].(float64))
					data0["do"] = int(data0["do"].(float64))
					res["player"] = data0
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
			}else if kk == "count" {
				trow["count"] = vv
			}else if kk == "before"{
				trow["before"] = vv
			}else if kk == "after" {
				trow["after"] = vv
			}else if kk == "reason" {
				trow["reason"] = vv
			}else if kk == "eid" {
				trow["eid"] = vv
			}else if kk == "time" {
				trow["time"] = time.Unix(int64(vv.(float64)), 0).Format("2006-01-02 15:04")
			}

		}
		res = append(res, trow)

	}


	return res
}