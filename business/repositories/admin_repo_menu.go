package repositories

import (
	"github.com/jinzhu/gorm"
	"gothic-admin/business/models"
	"strings"
)

type UpdateMenuRule func (user *models.Menu, params models.MenuForm)

func (r *userRepo) GetMenuByID(id uint) (menu models.Menu) {
	r.db.First(&menu, id)
	return
}

func (r *userRepo) GetMenuBy(params models.MenuForm) (menu models.Menu) {
	r.db.Where("tag = ? AND type = ? AND parent_id = ?", params.Tag, params.Type, params.ParentID).First(&menu)
	return
}


func (r *userRepo) AddMenu(params models.MenuForm) (menu models.Menu, ok bool) {

	menu = models.Menu{
		Tag:       params.Tag,
		Name:      params.Name,
		Type:      params.Type,
		ParentID:  params.ParentID,
	}

	r.db.Create(&menu)
	ok = !r.db.NewRecord(menu)

	if params.Type == models.MenuMenu {

		for method, name := range map[string]string {
			models.MethodGet:		"查询",
			models.MethodAdd: 		"创建",
			models.MethodUpdate:	"更新",
			models.MethodDelete:	"删除",
		} {
			auth := models.Menu{
				Tag:       method,
				Name:      name,
				Type:      models.MenuAuth,
				ParentID:  menu.ID,
			}
			r.db.Create(&auth)
		}

	}
	return
}

func (r *userRepo) UpdateMenu(params models.MenuForm, updateRule ...UpdateMenuRule) (menu models.Menu, err error) {
	menu = r.GetMenuByID(params.ID)
	if len(updateRule) > 0 {
		for _, update := range updateRule {
			update(&menu, params)
		}
	} else {
		// 默认更新方式
		menu.Name = params.Name
		menu.Tag = params.Tag
		menu.ParentID = params.ParentID
		if menu.Type == models.MenuMenu && len(params.Children) > 0 {
			for _, auth := range params.Children {
				if auth.ID == 0 {
					r.AddMenu(auth)
				} else {
					_, err = r.UpdateMenu(auth)
					if err != nil {
						return
					}
				}
			}
		}
	}
	err = r.db.Save(&menu).Error
	return
}

func (r *userRepo) DeleteMenu(ids []uint) error {
	return r.db.Where(ids).Delete(&models.Menu{}).Error
}

func (r *userRepo) GetMenuPage(current int, pageSize int, params map[string]string) (menus []models.Menu, count int) {
	menuSearch := func (params map[string]string) func (db *gorm.DB) *gorm.DB {
		return func (db *gorm.DB) *gorm.DB {
			if v, ok := params["type"]; ok && len(v) > 0 {
				vs := strings.Split(v, ",")
				db = db.Where("type IN (?)", vs)
			}
			return db
		}
	}

	r.db.Offset((current - 1) * pageSize).Limit(pageSize).Scopes(menuSearch(params)).Find(&menus).Count(&count)
	return
}


func (r *userRepo) FindMenu(tags []string, method string) (menu models.Menu, target models.Menu) {

	if len(tags) == 0 {
		menu = r.GetMenuBy(models.MenuForm{
			Tag:      tags[0],
			Type:     models.MenuMenu,
			ParentID: 0,
		})
	} else {
		for i, tag := range tags {
			t := models.MenuGroup
			if i == len(tags) - 1 {
				t = models.MenuMenu
			}
			var pid uint
			pid = 0
			if menu.ID > 0 {
				pid = menu.ID
			}
			menu = r.GetMenuBy(models.MenuForm{
				Tag:      tag,
				Type:     t,
				ParentID: pid,
			})
		}
	}

	if menu.ID > 0 {
		target = r.GetMenuBy(models.MenuForm{
			Tag:      method,
			Type:     models.MenuAuth,
			ParentID: menu.ID,
		})
	}
	return
}