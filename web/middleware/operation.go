package middleware

import (
	"fmt"
	"github.com/bbcloudGroup/gothic/di"
	"github.com/kataras/iris/v12"
	"gothic-admin/business/services"
)

func Operation(ctx iris.Context) {

	uid, _ := ctx.Params().GetUint(KeyUserID)
	menuID, _ := ctx.Params().GetUint(KeyMenuID)
	methodID, _ := ctx.Params().GetUint(KeyMethodID)

	di.Invoke(func(log services.LogService) {
		log.LogOperation(uid, menuID, methodID, fmt.Sprintln(ctx.Request()))
	})
}
