package repositories

import (
	"gothic-admin/business/models"
	"gothic-admin/datasource"
)


type AdminRepo interface {
	// user
	GetUserByID(id uint) models.User
	AddUser(params models.UserForm) (models.User, bool)
	UpdateUser(params models.UserForm, updateRule ...UpdateUserRule) (models.User, error)
	DeleteUser(ids []uint) error
	GetUserPage(current int, pageSize int, params map[string]string) ([]models.User, int)
	GetUserByMail(mail string) models.User
	GetUserByMobile(mobile string) models.User
	SetToken(user *models.User, token string)

	// role
	GetRoleByID(id uint) models.Role
	AddRole(params models.RoleForm) (models.Role, bool)
	UpdateRole(params models.RoleForm, updateRule ...UpdateRoleRule) (models.Role, error)
	DeleteRole(ids []uint) error
	GetRolePage(current int, pageSize int, params map[string]string) ([]models.Role, int)
	GetRoleByTag(tag string) models.Role
	GetRoleByName(nam string) models.Role

	// menu
	GetMenuByID(id uint) models.Menu
	AddMenu(params models.MenuForm) (models.Menu, bool)
	UpdateMenu(params models.MenuForm, updateRule ...UpdateMenuRule) (models.Menu, error)
	DeleteMenu(ids []uint) error
	GetMenuPage(current int, pageSize int, params map[string]string) ([]models.Menu, int)
	GetMenuBy(params models.MenuForm) models.Menu
	FindMenu(tags []string, method string) (models.Menu, models.Menu)

	// log
	AddLog(userID uint, menuID uint, methodID uint, data string) (models.Log, bool)
	DeleteLog(ids []uint) error
	GetLogPage(current int, pageSize int, params map[string]string) ([]models.Log, int)
}

func NewAdminRepo(db datasource.Admin) AdminRepo {
	return &userRepo{db:db}
}

type userRepo struct {
	db datasource.Admin
}













