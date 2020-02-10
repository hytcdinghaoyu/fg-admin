package models

import (
	"fmt"
	"time"
)

type OperationLog struct {
	ID         uint `gorm:"primary_key"`
	CreatedAt  time.Time
	Uid        int
	Method     string `gorm:"VARCHAR(10)"`
	Path       string `gorm:"VARCHAR(50)"`
	Params     string `gorm:"VARCHAR(200)"`
	IP         string `gorm:"VARCHAR(50)"`
	StatusCode int
}

// 分页获取用户记录
func GetAllOpLogs(name, orderBy string, page, limit int) (logs []*OperationLog) {
	if err := GetAll(name, orderBy, page, limit).Find(&logs).Error; err != nil {
		fmt.Printf("GetAllOpLogsErr:%s", err)
	}
	return
}

// 总用户数
func GetLogsCount() int {
	var count int
	DB.Model(&OperationLog{}).Count(&count)
	return count
}
