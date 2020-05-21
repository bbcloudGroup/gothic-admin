package providers

import (
	"github.com/bbcloudGroup/gothic/di"
	"gothic-admin/business/services"
)

func RegisterServices(container di.Container) {
	container.Register(services.NewUserService)
	container.Register(services.NewCaptchaService)
	container.Register(services.NewRoleService)
	container.Register(services.NewMenuService)
	container.Register(services.NewLogService)
}
