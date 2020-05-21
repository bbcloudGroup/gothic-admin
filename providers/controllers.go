package providers

import (
	"github.com/bbcloudGroup/gothic/di"
	"gothic-admin/web/controllers/admin"
)

func RegisterController(container di.Container) {
	container.Register(admin.NewAdminController)
	container.Register(admin.NewApiController)
}
