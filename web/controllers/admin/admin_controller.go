package admin

import (
	"github.com/kataras/iris/v12/mvc"
	"gothic-admin/business/services"
)

type AdminController struct {
	User services.UserService
	Role services.RoleService
	Menu services.MenuService
	Log	 services.LogService
}

func NewAdminController(
	user services.UserService,
	role services.RoleService,
	menu services.MenuService,
	log  services.LogService) AdminController {
	return AdminController{User: user, Role: role, Menu: menu, Log: log}
}

func (c *AdminController) BeforeActivation(b mvc.BeforeActivation) {


}


