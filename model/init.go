package models

import (
	"bytes"
	"encoding/json"
	"fg-admin/config"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/patrickmn/go-cache"
	"github.com/pelletier/go-toml"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	DB       = New()
	MemCache = cache.New(10*time.Minute, 1*time.Hour)
)

func PostLogin(path string, body []byte) map[string]interface{} {
	url := config.LoginServerAddr + path
	rsp, err := http.Post(url, "application/json", bytes.NewReader(body))
	ret := make(map[string]interface{})
	if err != nil {
		fmt.Println(err)
		return ret
	}

	defer rsp.Body.Close()
	rspBody, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		fmt.Println(err)
	}


	err = json.Unmarshal(rspBody, &ret)
	if err != nil {
		fmt.Println(err)
	}
	return ret

}

func PostGm(path string, body []byte) map[string]interface{} {
	url := config.GmServerAddr + path
	rsp, err := http.Post(url, "application/json", bytes.NewReader(body))
	ret := make(map[string]interface{})
	if err != nil {
		fmt.Println(err)
		return ret
	}

	defer rsp.Body.Close()
	rspBody, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(rspBody, &ret)
	if err != nil {
		fmt.Println(err)
	}
	return ret

}

func PostGmArr(path string, body []byte) []interface{} {
	url := config.GmServerAddr + path
	rsp, err := http.Post(url, "application/json", bytes.NewReader(body))
	ret := []interface{}{}
	if err != nil {
		fmt.Println("http post err", err)
		return ret
	}

	defer rsp.Body.Close()
	rspBody, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		fmt.Println("read err", err)
	}

	err = json.Unmarshal(rspBody, &ret)
	if err != nil {
		fmt.Println("json decode err", err)
	}
	return ret
}

func PostLog(path string, body []byte) map[string]interface{} {
	url := config.LogServerAddr + path
	// body = []byte(`{"btime":"2020-03-21 18:22:47","etime":"2020-03-23 18:22:47", "eid":3001, "uid":2434880, "skip":0, "limit":20, "sid": 1}`)
	rsp, err := http.Post(url, "application/json", bytes.NewReader(body))
	ret := map[string]interface{}{}
	if err != nil {
		fmt.Println("http post err", err)
		return ret
	}

	defer rsp.Body.Close()
	rspBody, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		fmt.Println("read err", err)
	}

	err = json.Unmarshal(rspBody, &ret)
	if err != nil {
		fmt.Println("json decode err", err)
	}
	return ret
}



func PostGmStruct(path string, req interface{}, ret interface{}){

	body, err := json.Marshal(req)
	if err != nil {
		fmt.Println("json marshal err", err)
	}

	url := config.GmServerAddr + path
	rsp, err := http.Post(url, "application/json", bytes.NewReader(body))
	if err != nil {
		fmt.Println("post gm struct err", err)
	}

	defer rsp.Body.Close()
	rspBody, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		fmt.Println("read err", err)
	}

	err = json.Unmarshal(rspBody, ret)
	if err != nil {
		fmt.Println(err)
	}

}

// New setup DB connection
func New() *gorm.DB {

	driver := config.Conf.Get("database.driver").(string)
	configTree := config.Conf.Get(driver).(*toml.Tree)
	userName := configTree.Get("username").(string)
	password := configTree.Get("password").(string)
	databaseName := configTree.Get("db").(string)
	connect := userName + ":" + password + "@/" + databaseName + "?charset=utf8&parseTime=True&loc=Local"

	DB, err := gorm.Open(driver, connect)

	if err != nil {
		panic(fmt.Sprintf("failed to connect to database, got err=%+v", err))
	}

	return DB
}

func GetAll(string, orderBy string, page, limit int) *gorm.DB {
	TDB := DB
	if len(orderBy) > 0 {
		TDB = TDB.Order(orderBy + "desc")
	} else {
		TDB = TDB.Order("created_at desc")
	}

	if len(string) > 0 {
		TDB = TDB.Where("name LIKE  ?", "%"+string+"%")
	}

	if page > 0 {
		TDB = TDB.Offset((page - 1) * limit)
	}

	if limit > 0 {
		TDB = TDB.Limit(limit)
	}

	return TDB
}
