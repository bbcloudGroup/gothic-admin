package repositories

import (
	"github.com/jinzhu/gorm"
	"gothic-admin/business/models"
	"strings"
)

func (r *userRepo) AddLog(userID uint, menuID uint, methodID uint, data string) (log models.Log, ok bool) {
	log = models.Log{
		UserID:    userID,
		MenuID:    menuID,
		MethodID:  methodID,
		Data:      data,
	}

	r.db.Create(&log)
	ok = !r.db.NewRecord(log)
	return
}

func (r *userRepo) DeleteLog(ids []uint) error {
	return r.db.Where(ids).Delete(&models.Log{}).Error
}

func (r *userRepo) GetLogPage(current int, pageSize int, params map[string]string) (logs []models.Log, count int) {
	var log models.Log
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
			if v, ok := params["menu"]; ok && len(v) > 0 {
				db = db.Where("menu_id = ? OR method_id = ?", v, v)
			}
			if v, ok := params["user"]; ok && len(v) > 0 {
				db = db.Where("user_id = ?", v)
			}
			return db
		}
	}

	r.db.Offset((current - 1) * pageSize).Limit(pageSize).Scopes(roleSearch(params)).
		Model(&log).Preload("Menu").Preload("Method").Preload("User").Order("created_at desc").Find(&logs).Count(&count)
	return
}

