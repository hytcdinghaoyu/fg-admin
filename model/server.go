package models

import (
	"fg-admin/config"
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

var (
	MServerType = map[int32]string{
		SERVER_TYPE_GATE:   "网关服",
		SERVER_TYPE_LOGIC:  "逻辑服",
		SERVER_TYPE_BATTLE: "战斗服",
		SERVER_TYPE_MGR:    "管理服",
		SERVER_TYPE_LOG:    "日志服",
		SERVER_TYPE_GM:     "GM服",
		SERVER_TYPE_LOGIN:  "登录服",
	}
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

func GetServerStatus(tp string) interface{} {
	ret := HttpPost(config.GmServerAddr+"/server/serverlist", []byte(``))
	serverMap := GetServerMap()
	mLogic := []map[string]interface{}{}
	mBattle := []map[string]interface{}{}
	if data, ok := ret["data"]; ok {
		if dataMap, ok := data.([]interface{}); ok {
			for _, v := range dataMap {
				if m, ok := v.(map[string]interface{}); ok {

					mt := m
					mt["ServerType"] = GetServerName(int32(mt["Stype"].(float64)), int32(mt["Sid"].(float64)))
					mt["StartTime"] = time.Unix(int64(mt["StartTime"].(float64)), 0).Format("2006-01-02 15:04")

					if nm, ok := serverMap[int32(mt["Sid"].(float64))]; ok {
						mt["ServerName"] = nm
					}

					if t, ok := mt["Time"].(float64); ok {
						if t > 0 {
							mt["Status"] = "开启"
						} else {
							mt["Status"] = "关闭"
						}
					} else {
						mt["Status"] = "关闭"
					}

					sec := time.Second * time.Duration(mt["RunningTime"].(float64))
					mt["RunningTime"] = sec.String()

					if connList, ok := mt["ConnList"].([]interface{}); ok {
						ConnListName := []string{}
						for _, vv := range connList {
							if mm, ok := vv.(map[string]interface{}); ok {
								conName := GetServerName(int32(mm["Stype"].(float64)), int32(mm["Sid"].(float64)))
								ConnListName = append(ConnListName, conName)
							}
						}
						mt["ConnListName"] = ConnListName
					}

					stype := int(m["Stype"].(float64))
					if stype == SERVER_TYPE_LOGIC {
						mLogic = append(mLogic, mt)
					} else if stype == SERVER_TYPE_BATTLE {
						mBattle = append(mBattle, mt)
					}

				}
			}
		}

		if tp == "logic" {
			return mLogic
		} else if tp == "battle" {
			return mBattle
		}
	}
	return nil
}
