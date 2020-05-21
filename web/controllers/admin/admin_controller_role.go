package admin

import (
	"errors"
	"github.com/kataras/iris/v12"
	"gothic-admin/business/models"
)

func (c *AdminController) GetRole(ctx iris.Context) models.Page {
	current, _ := ctx.URLParamInt("current")
	pageSize, _ := ctx.URLParamInt("pageSize")
	return c.Role.GetPage(current, pageSize, ctx.URLParams())
}

func (c *AdminController) PostRole(ctx iris.Context) models.Response {

	var params models.RoleForm
	_ = ctx.ReadJSON(&params)

	msg := ""
	err := errors.New("未授权操作")
	switch params.Method {
	case models.MethodAdd:
		_, err = c.Role.Add(params)
	case models.MethodUpdate:
		_, err = c.Role.Update(params)
	case models.MethodDelete:
		err = c.Role.Delete(params)
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
