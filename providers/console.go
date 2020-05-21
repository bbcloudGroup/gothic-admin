package providers

import (
	"github.com/bbcloudGroup/gothic/di"
	"gothic-admin/boot/console"
)

func RegisterConsole(container di.Container) {

	container.Register(console.NewPrinter)

	container.Register(console.NewAdminMigrate)

}
