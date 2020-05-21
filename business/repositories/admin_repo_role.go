package repositories

import (
	"github.com/jinzhu/gorm"
	"gothic-admin/business/models"
	"strings"
)

type UpdateRoleRule func (user *models.Role, params models.RoleForm)

func (r *userRepo) GetRoleByID(id uint) (role models.Role) {
	r.db.First(&role, id)
	return
}


func (r *userRepo) AddRole(params models.RoleForm) (role models.Role, ok bool) {
	role = models.Role{
		Name:	params.Name,
		Tag:	params.Tag,
	}
	r.db.Create(&role)
	ok = !r.db.NewRecord(role)
	return
}

func (r *userRepo) UpdateRole(params models.RoleForm, updateRule ...UpdateRoleRule) (role models.Role, err error) {
	role = r.GetRoleByID(params.ID)
	if len(updateRule) > 0 {
		for _, update := range updateRule {
			update(&role, params)
		}
	} else {
		// 默认更新方式
		role.Name = params.Name
		var menus []models.Menu
		r.db.Where(params.Menu).Find(&menus)
		r.db.Model(&role).Association("Menu").Replace(menus)
	}
	err = r.db.Save(&role).Error
	return
}

func (r *userRepo) DeleteRole(ids []uint) error {
	return r.db.Where(ids).Delete(&models.Role{}).Error
}




func (r *userRepo) GetRolePage(current int, pageSize int, params map[string]string) (roles []models.Role, count int) {

	var role models.Role
	roleSearch := func (params map[string]string) func (db *gorm.DB) *gorm.DB {
		return func (db *gorm.DB) *gorm.DB {
			if v, ok := params["id"]; ok && len(v) > 0 {
				vs := strings.Split(v, ",")
				if len(vs) > 1 {
					db = db.Where("id IN (?)", vs)
				} else {
					db = db.Where("id = ?", v)
				}
			}
			if v, ok := params["name"]; ok && len(v) > 0 {
				db = db.Where("name LIKE ?", "%" + v + "%")
			}
			if v, ok := params["tag"]; ok && len(v) > 0 {
				db = db.Where("tag LIKE ?", "%" + v + "%")
			}
			return db
		}
	}

	r.db.Offset((current - 1) * pageSize).Limit(pageSize).Scopes(roleSearch(params)).
		Model(&role).Preload("Menu").Find(&roles).Count(&count)
	return
}


func (r *userRepo) GetRoleByTag(tag string) (role models.Role) {
	r.db.Where("tag = ?", tag).First(&role)
	return
}

func (r *userRepo) GetRoleByName(name string) (role models.Role) {
	r.db.Where("name = ?", name).First(&role)
	return
}

