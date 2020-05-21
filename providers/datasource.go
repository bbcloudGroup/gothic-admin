package providers

import (
	"github.com/bbcloudGroup/gothic/di"
	"gothic-admin/datasource"
)

func RegisterDatabase(container di.Container) {

	container.Register(datasource.NewAdmin)
	container.Register(datasource.NewCache)
}

