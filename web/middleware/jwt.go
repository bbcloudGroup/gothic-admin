package middleware

import (
	"github.com/kataras/iris/v12"
	"gothic-admin/business/models"
	"gothic-admin/utils"
	"strconv"
	"strings"
)

func Jwt(ctx iris.Context) {

	auth := ctx.GetHeader("Authorization")
	if len(auth) > 0 {
		claims, ok := utils.JwtValidate(strings.Split(auth, " ")[1])
		if ok {
			ctx.Params().Set(KeyUserID, strconv.Itoa(int(claims.UserID)))
			ctx.Next()
			return
		}
	}

	_, _ = ctx.JSON(models.Response{
		Status:  "error",
		Message: "登录过期，请重新登录",
	})

}