package middleware

import (
	"encoding/json"
	"fg-admin/constant"
	models "fg-admin/model"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"strconv"
	"strings"
	"time"
)

var excludePath = []string{"list", "index", "tree", "login", "logout"}

func NewOperateLogger() iris.Handler {

	c := logger.Config{}

	c.AddSkipper(func(ctx iris.Context) bool {
		path := ctx.Path()
		for _, ext := range excludePath {
			if strings.Contains(path, ext) {
				return true
			}
		}
		return false
	})

	c.LogFuncCtx = func(ctx iris.Context, latency time.Duration) {

		if ctx.Method() != "POST" {
			return
		}

		uid := ctx.Values().GetString(constant.CTX_UID)
		req := ctx.FormValues()

		oplog := new(models.OperationLog)
		oplog.StatusCode = ctx.GetStatusCode()
		oplog.Uid, _ = strconv.Atoi(uid)
		oplog.Path = ctx.Path()
		oplog.Method = ctx.Method()
		oplog.IP = ctx.RemoteAddr()
		paramBytes ,err := json.Marshal(req)
		if err == nil {
			oplog.Params = string(paramBytes)
		}

		if err := models.DB.Create(oplog).Error; err != nil {
			fmt.Printf("CreateOpLogErr:%s", err)
		}
	}

	return logger.New(c)
}
