package repositories

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"gothic-admin/business/models"
	"gothic-admin/utils"
	"strconv"
	"strings"
)


type UpdateUserRule func (user *models.User, params models.UserForm)

func (r *userRepo) AddUser(params models.UserForm) (user models.User, ok bool) {

	user = models.User{
		Password: utils.Md5s(params.Password),
		Name:     params.Mail,
		Avatar:   fmt.Sprintf("https://secure.gravatar.com/avatar/%s?d=identicon", utils.Md5s(params.Mail)),
		Mail:     params.Mail,
		Mobile:   params.Mobile,
		Status:	  true,
	}

	r.db.Create(&user)

	role := r.GetRoleByTag("guest")

	r.db.Model(&user).Association("Role").Append(role)
	ok = !r.db.NewRecord(user)
	return
}


func (r *userRepo) UpdateUser(params models.UserForm, updateRule ...UpdateUserRule) (user models.User, err error) {

	user = r.GetUserByID(params.ID)
	if len(updateRule) > 0 {
		for _, update := range updateRule {
			update(&user, params)
		}
	} else {
		// 默认更新方式
		user.Name = params.Name

		var roles []models.Role
		var rids []uint
		for _, role := range params.Role {
			rids = append(rids, role.ID)
		}
		r.db.Where(rids).Find(&roles)
		r.db.Model(&user).Association("Role").Replace(roles)
	}

	err = r.db.Save(&user).Error
	return
}


func (r *userRepo) DeleteUser(ids []uint) error {
	return r.db.Where(ids).Delete(&models.User{}).Error
}


func (r *userRepo) GetUserPage(current int, pageSize int, params map[string]string) (users []models.User, count int) {

	var user models.User
	userSearch := func (params map[string]string) func (db *gorm.DB) *gorm.DB {
		return func (db *gorm.DB) *gorm.DB {

			if v, ok := params["id"]; ok && len(v) > 0 {
				db = db.Where("id = ?", v)
			}
			if v, ok := params["name"]; ok && len(v) > 0 {
				db = db.Where("name LIKE ?", "%" + v + "%")
			}
			if v, ok := params["mobile"]; ok && len(v) > 0 {
				db = db.Where("mobile = ?", v)
			}
			if v, ok := params["mail"]; ok && len(v) > 0 {
				db = db.Where("mail = ?", v)
			}
			if v, ok := params["status"]; ok && len(v) > 0 {
				s, _ := strconv.ParseBool(v)
				db = db.Where("status = ?", s)
			}
			if v, ok := params["created_at"]; ok && len(v) > 0 {
				vs := strings.Split(v, ",")
				db = db.Where("created_at BETWEEN ? AND ?", vs[0], vs[1])
			}
			if v, ok := params["roles"]; ok && len(v) > 0 {
				vs := strings.Split(v, ",")
				db = db.Joins("JOIN user_role ON user.id = user_role.user_id AND user_role.role_id IN (?)", vs)
			}
			return db
		}
	}
	r.db.Offset((current - 1) * pageSize).Limit(pageSize).Scopes(userSearch(params)).
		Model(&user).Preload("Role").Find(&users).Count(&count)
	return
}





func (r *userRepo) GetUserByMobile(mobile string) models.User {
	var user models.User
	r.db.Preload("Role").Where("mobile = ?", mobile).First(&user)
	return user
}

func (r *userRepo) GetUserByID(id uint) models.User {
	var user models.User
	r.db.Preload("Role").First(&user, id)
	return user
}

func (r *userRepo) GetUserByMail(mail string) models.User {
	var user models.User
	r.db.Preload("Role").Where("mail = ?", mail).First(&user)
	return user
}

func (r *userRepo) SetToken(user *models.User, token string) {
	user.RememberToken = token
	r.db.Save(user)
}
