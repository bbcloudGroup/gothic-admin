package providers

import (
	"github.com/bbcloudGroup/gothic/di"
	"gothic-admin/business/repositories"
)

func RegisterRepositories(container di.Container) {

	container.Register(repositories.NewCaptchaRepo)
	container.Register(repositories.NewAdminRepo)
}