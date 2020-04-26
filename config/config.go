package config

import (
	"encoding/csv"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/pelletier/go-toml"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
)

var (
	Conf            = New()
	ItemConf        = NewItemConf()
	ItemConfMap     = make(map[int]string)
	HeroConfMap     = NewConfMap("HeroMain", 23)
	EquipConfMap    = NewConfMap("Equip", 1)
	LoginServerAddr = Conf.Get("login_server.host").(string)
	GmServerAddr    = Conf.Get("gm_server.host").(string)
	LogServerAddr   = Conf.Get("log_server.host").(string)
)

type JwtClaims struct {
	*jwt.StandardClaims
	UserInfo
}

type UserInfo struct {
	Uid      uint
	Username string
	RoleId   uint
	AuthStr  string
}

// 道具配表
type Item struct {
	ID   int
	Name string
}

/**
 * 返回单例实例
 * @method New
 */
func New() *toml.Tree {
	config, err := toml.LoadFile("./config.toml")

	if err != nil {
		fmt.Println("TomlError ", err.Error())
	}

	return config
}

func NewItemConf() []Item {
	dat, err := ioutil.ReadFile("./upload/Item.csv")
	if err != nil {
		panic(err)
	}
	r := csv.NewReader(strings.NewReader(string(dat[:])))

	var sliItem []Item
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		id, err := strconv.Atoi(record[0])
		if id == 0 {
			continue
		}

		if err != nil {
			panic(err)
		}

		i := Item{ID: id, Name: record[1]}
		sliItem = append(sliItem, i)

		ItemConfMap[id] = record[1]
	}

	return sliItem
}

func NewConfMap(name string, offset int) map[int]string {

	path := fmt.Sprintf("./upload/%s.csv", name)
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	r := csv.NewReader(strings.NewReader(string(dat[:])))

	m := map[int]string{}
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		id, err := strconv.Atoi(record[0])
		if id == 0 {
			continue
		}

		if err != nil {
			panic(err)
		}

		m[id] = record[offset]
	}

	return m
}
