package admin


import (
	"errors"
	"github.com/kataras/iris/v12"
	"gothic-admin/business/models"
)

func (c *AdminController) GetMenu(ctx iris.Context) models.Page {
	//current, _ := ctx.URLParamInt("current")
	//pageSize, _ := ctx.URLParamInt("pageSize")
	return c.Menu.GetPage(0, -1, ctx.URLParams())
}

func (c *AdminController) PostMenu(ctx iris.Context) models.Response {

	var params models.MenuForm
	_ = ctx.ReadJSON(&params)

	msg := ""
	err := errors.New("未授权操作")
	switch params.Method {
	case models.MethodAdd:
		_, err = c.Menu.Add(params)
	case models.MethodUpdate:
		_, err = c.Menu.Update(params)
	case models.MethodDelete:
		err = c.Menu.Delete(params)
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
