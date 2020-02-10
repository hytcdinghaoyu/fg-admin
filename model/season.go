package models

import "fmt"

const (
	ACTIVITY_ID_ASSIST = 1
)

var ActivityMap = map[int]string{ACTIVITY_ID_ASSIST: "英雄助战"}

type Season struct {
	ID         uint `gorm:"primary_key"`
	ServerId   int
	ActivityId int
	StartAt    string `gorm:"VARCHAR(50)"`
	EndAt      string `gorm:"VARCHAR(50)"`
}

func UpsertSeason(season Season) {

	s := Season{}
	if err := DB.Where("activity_id = ? and server_id = ?", season.ActivityId, season.ServerId).First(&s).Error; err == nil {
		s.StartAt = season.StartAt
		s.EndAt = season.EndAt
		DB.Save(s)
	}else{
		if err := DB.Create(&season).Error; err != nil {
			fmt.Printf("CreateSeasonErr:%s", err)
		}
	}

	return
}

func GetAllSeasons(name, orderBy string, offset, limit int) (seasons []Season) {

	if err := DB.Find(&seasons).Error; err != nil {
		fmt.Printf("GetAllSeasonsErr:%s", err)
	}
	return
}

// 通过 id 获取 user 记录
func GetSeasonById(id uint) (s Season, err error) {
	season := new(Season)
	season.ID = id

	fmt.Println(season)

	if err := DB.First(season).Error; err != nil {
		fmt.Printf("GetSeasonByIdErr:%s", err)
	}

	return *season, nil
}
