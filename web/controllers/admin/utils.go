package admin

import (
	"github.com/kataras/iris/v12"
	"gothic-admin/web/middleware"
)

func getUserID(ctx iris.Context) uint {
	userID, _ := ctx.Params().GetUint(middleware.KeyUserID)
	return userID
}
