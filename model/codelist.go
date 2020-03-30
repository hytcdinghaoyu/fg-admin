package models

import "encoding/json"

type CodeList struct {
	Id             int32         `json:"id" bson:"id"`
	Name           string        `json:"name" bson:"name"`     // 名称
	GenNum         int32         `json:"gnum" bson:"gnum"`     // 生成数量
	UsedNum        int32         `json:"unum" bson:"unum"`     // 使用数量
	CreateTime     int64         `json:"ct" bson:"ct"`         // 添加时间
	ExpirationDate int64         `json:"et" bson:"et"`         // 过期时间
}


func DeleteCode(id int) (status float64, msg string) {

	req := make(map[string]interface{})
	req["id"] = id
	req["act"] = "delcode"

	body, err := json.Marshal(req)
	if err != nil {
		return STATUS_OP_FAILED, MSG_OP_FIALED
	}

	ret := PostGm("/", body)
	if val, ok := ret["code"]; ok {
		return val.(float64), ret["msg"].(string)
	}

	return STATUS_OP_FAILED, MSG_OP_FIALED
}