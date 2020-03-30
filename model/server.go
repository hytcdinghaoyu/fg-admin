package models

import (
	"strconv"
	"time"
)

const (
	SERVER_TYPE_GATE   = 1 // 网关服
	SERVER_TYPE_LOGIC  = 2 // 逻辑服
	SERVER_TYPE_BATTLE = 3 // 战斗服
	SERVER_TYPE_MGR    = 4 // 管理服
	SERVER_TYPE_LOG    = 5 // 日志服
	SERVER_TYPE_GM     = 6 // GM服
	SERVER_TYPE_LOGIN  = 7 // 登陆服
)

func GetServerName(stype, sid int32) string {
	switch stype {
	case SERVER_TYPE_GATE:
		return "网关[" + strconv.Itoa(int(sid)) + "]"
	case SERVER_TYPE_LOGIC:
		return "逻辑服[" + strconv.Itoa(int(sid)) + "]"
	case SERVER_TYPE_BATTLE:
		return "战斗服[" + strconv.Itoa(int(sid)) + "]"
	case SERVER_TYPE_MGR:
		return "管理服[" + strconv.Itoa(int(sid)) + "]"
	case SERVER_TYPE_LOG:
		return "日志服[" + strconv.Itoa(int(sid)) + "]"
	case SERVER_TYPE_GM:
		return "GM服[" + strconv.Itoa(int(sid)) + "]"
	case SERVER_TYPE_LOGIN:
		return "登陆服[" + strconv.Itoa(int(sid)) + "]"
	}
	return ""
}

func GetServerStatus() interface{} {
	ret := PostGm("/server/serverlist", []byte(``))
	serverMap := GetServerMap()
	if data, ok := ret["data"]; ok {
		if dataMap, ok := data.([]interface{}); ok {
			for _, v := range dataMap {
				if m, ok := v.(map[string]interface{}); ok {
					m["ServerType"] = GetServerName(int32(m["Stype"].(float64)), int32(m["Sid"].(float64)))
					m["StartTime"] = time.Unix(int64(m["StartTime"].(float64)), 0).Format("2006-01-02 15:04")

					if nm, ok := serverMap[int32(m["Sid"].(float64))]; ok {
						m["ServerName"] = nm
					}

					if m["Time"].(float64) > 0 {
						m["Status"] = "开启"
					}else {
						m["Status"] = "关闭"
					}

					sec := time.Second * time.Duration(m["RunningTime"].(float64))
					m["RunningTime"] = sec.String()

					if connList, ok := m["ConnList"].([]interface{}); ok {
						ConnListName := []string{}
						for _, vv := range connList {
							if mm, ok := vv.(map[string]interface{}); ok {
								conName := GetServerName(int32(mm["Stype"].(float64)), int32(mm["Sid"].(float64)))
								ConnListName = append(ConnListName, conName)
							}
						}
						m["ConnListName"] = ConnListName
					}

				}
			}
		}

		return data
	}
	return nil
}