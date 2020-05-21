package admin

import (
	"errors"
	"github.com/kataras/iris/v12"
	"gothic-admin/business/models"
)

func (c *AdminController) GetUser(ctx iris.Context) models.Page {
	current, _ := ctx.URLParamInt("current")
	pageSize, _ := ctx.URLParamInt("pageSize")
	return c.User.GetPage(current, pageSize, ctx.URLParams())
}

func (c *AdminController) PostUser(ctx iris.Context) models.Response {

	var params models.UserForm
	_ = ctx.ReadJSON(&params)

	msg := ""
	err := errors.New("未授权操作")
	switch params.Method {
	case models.MethodAdd:
		_, err = c.User.Add(params)
	case models.MethodUpdate:
		_, err = c.User.Update(params)
	case models.MethodDelete:
		err = c.User.Delete(params)
	case "resetPassword":
		_, err = c.User.ResetPassword(params)
	case "statusChange":
		_, err = c.User.StatusChange(params)
	case "approval":
		// 批量操作
	}

	if err != nil {
		return models.Response{
			Status:  "error",
			Message: err.Error(),
		}
	} else {
		return models.Response{
			Status:  "ok",
			Message: msg,
		}
	}
}
