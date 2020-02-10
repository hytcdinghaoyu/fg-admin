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
	Conf     = New()
	ItemConf = NewItemConf()
	LoginServerAddr = Conf.Get("login_server.host").(string)
	GmServerAddr = Conf.Get("gm_server.host").(string)
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

type Item struct {
	ID   int
	Name string
}

/**
 * 返回单例实例
 * @method New
 */
func New() *toml.Tree {
	config, err := toml.LoadFile("./config/config.toml")

	if err != nil {
		fmt.Println("TomlError ", err.Error())
	}

	return config
}

func NewItemConf() []Item {
	dat, err := ioutil.ReadFile("./upload/item.csv")
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
	}

	return sliItem
}
