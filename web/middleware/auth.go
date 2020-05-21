package middleware

import (
	"github.com/bbcloudGroup/gothic/di"
	"github.com/kataras/iris/v12"
	"gothic-admin/business/models"
	"gothic-admin/business/services"
	"strconv"
	"strings"
)


const (
	KeyUserID = "_admin_user_id"
	KeyMenuID = "_admin_menu_id"
	KeyMethodID = "_admin_method_id"
)


func Auth(ctx iris.Context) {

	method := models.MethodGet
	if ctx.Method() == iris.MethodPost {
		form := models.Form{}
		err := ctx.ReadJSON(&form)
		if err != nil {
			method = ""
		} else {
			method = form.Method
		}
	}

	b := false
	var menu models.Menu
	var target models.Menu
	if len(method) > 0 {
		di.Invoke(func(menuService services.MenuService) {
			p := strings.Split(ctx.GetCurrentRoute().Path(), "/")
			menu, target = menuService.FindMenu(p[2:], method)
		})
	}
	if target.ID > 0 {
		uid, _ := ctx.Params().GetUint(KeyUserID)

		ctx.Params().Set(KeyMenuID, strconv.Itoa(int(menu.ID)))
		ctx.Params().Set(KeyMethodID, strconv.Itoa(int(target.ID)))

		if uid == 1 {
			ctx.Next()
			return
		}
		di.Invoke(func (userService services.UserService) {
			b = userService.CheckMenu(uid, target)
		})
	}

	if b {
		ctx.Next()
		return
	}

	_, _ = ctx.JSON(models.Response{
		Status:  "error",
		Message: "用户未被授权该操作",
	})


}

