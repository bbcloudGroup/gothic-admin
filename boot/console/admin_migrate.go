package console

import (
	"gothic-admin/business/models"
	"gothic-admin/datasource"
)

type AdminMigrate struct {
	admin datasource.Admin
}

func NewAdminMigrate(admin datasource.Admin) AdminMigrate {
	return AdminMigrate{admin:admin}
}

func (a AdminMigrate) Run() {
	a.admin.AutoMigrate(
		&models.Role{},
		&models.User{},
		&models.Menu{},
		&models.Log{})
}
