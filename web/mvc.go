package web

import (
	"github.com/bbcloudGroup/gothic/di"
	"github.com/bbcloudGroup/gothic/routes"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"gothic-admin/web/controllers/admin"
	"gothic-admin/web/middleware"
)


func Mvc(app *iris.Application) map[string]routes.Route {

	return map[string]routes.Route {
		"/admin/api": routes.New(Api),
		"/admin/admin": routes.New(Admin).WithMiddleware(middleware.Jwt, middleware.Auth).WithTerminate(middleware.Operation),
	}
}


func Api(m *mvc.Application) {
	di.Invoke(func(controller admin.ApiController) {m.Handle(&controller)})
}

func Admin(m *mvc.Application) {
	di.Invoke(func(controller admin.AdminController) {m.Handle(&controller)})
}

