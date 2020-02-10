package middleware

import (
	"fg-admin/config"
	"fg-admin/constant"
	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris/v12/context"
	"strings"
)

// New returns a new auth middleware,
func NewAuth() context.Handler {
	return func(ctx context.Context) {
		auth := ctx.GetCookie("token")
		if auth == "" && ctx.Path() != "/user/login" {
			ctx.Redirect("/user/login")
			ctx.StopExecution()
		}

		// Validate token
		token := new(jwt.Token)
		var err error
		token, err = jwt.ParseWithClaims(auth, &config.JwtClaims{}, func(token *jwt.Token) (i interface{}, e error) {
			return []byte("secret"), nil
		})
		if (err != nil || !token.Valid || token == nil) && ctx.Path() != "/user/login" {
			ctx.Redirect("/user/login")
			ctx.StopExecution()
		}

		// Store user information from token into context.
		if token != nil {
			myClaims := token.Claims.(*config.JwtClaims)
			ctx.Values().Set(constant.CTX_UID, myClaims.Uid)
			ctx.Values().Set(constant.CTX_USERNAME, myClaims.Username)
			ctx.Values().Set(constant.CTX_ROLEID, myClaims.RoleId)

			if !strings.Contains("/user/login/user/logout/home/index", ctx.Path()) && !strings.Contains(myClaims.AuthStr, ctx.Path()) {
				ret := make(map[string]interface{})
				ret["status"] = -1
				ret["message"] = "Not Authorized"
				ctx.JSON(ret)
				return
			}

		}

		ctx.Next()
	}
}
