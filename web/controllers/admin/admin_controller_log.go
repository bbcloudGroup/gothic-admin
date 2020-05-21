package admin

import (
	"errors"
	"github.com/kataras/iris/v12"
	"gothic-admin/business/models"
)

func (c *AdminController) GetLog(ctx iris.Context) models.Page {
	current, _ := ctx.URLParamInt("current")
	pageSize, _ := ctx.URLParamInt("pageSize")
	return c.Log.GetPage(current, pageSize, ctx.URLParams())
}

func (c *AdminController) PostLog(ctx iris.Context) models.Response {

	var params models.Form
	_ = ctx.ReadJSON(&params)

	msg := ""
	err := errors.New("未授权操作")
	switch params.Method {
	case models.MethodDelete:
		err = c.Log.Delete(params)
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
