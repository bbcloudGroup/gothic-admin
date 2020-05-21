package services

import (
	"errors"
	"gothic-admin/business/models"
	"gothic-admin/business/repositories"
)

type RoleService interface {
	Add(params models.RoleForm) (models.Role, error)
	Update(params models.RoleForm) (models.Role, error)
	Delete(params models.RoleForm) error
	GetPage(current int, pageSize int, params map[string]string) models.Page
	GetMenuRoles() map[string][]string
}

func NewRoleService(repo repositories.AdminRepo) RoleService {
	return &roleService{repo:repo}
}

type roleService struct {
	repo repositories.AdminRepo
}

func (s *roleService) Add(params models.RoleForm) (role models.Role, err error) {
	if !(s.repo.GetRoleByTag(params.Tag).ID == 0) {
		err = errors.New("该角色标签已存在了")
		return
	}

	if !(s.repo.GetRoleByName(params.Name).ID == 0) {
		err = errors.New("该角色名称已存在了")
		return
	}

	role, ok := s.repo.AddRole(params)
	if !ok {
		err = errors.New("角色添加失败")
	}
	return
}

func (s *roleService) Update(params models.RoleForm) (models.Role, error) {
	return s.repo.UpdateRole(params)
}

func (s *roleService) Delete(params models.RoleForm) (err error) {
	if len(params.IDS) == 0 {
		return
	}

	for _, id := range params.IDS {
		if id == 1 {
			err = errors.New("无法删除超级管理员")
			return
		}
	}

	return s.repo.DeleteRole(params.IDS)
}


func (s *roleService) GetPage(current int, pageSize int, params map[string]string) models.Page {
	var data []interface{}

	roles, count :=  s.repo.GetRolePage(current, pageSize, params)

	for _, user := range roles {
		data = append(data, user)
	}

	return models.Page{
		Success:  true,
		Current:  current,
		PageSize: pageSize,
		Total:    count,
		Data:     data,
	}
}



func menuPath(res []string, menus []models.Menu, parent_id uint) []string {
	for _, menu := range menus {
		if menu.ID == parent_id {
			res = append(res, menu.Tag)
			res= menuPath(res, menus, menu.ParentID)
			return res
		}
	}
	return res
}



func (s *roleService) GetMenuRoles() (res map[string][]string) {

	res = make(map[string][]string)

	menus, _ := s.repo.GetMenuPage(0, -1, map[string]string{"type": "1,2"})
	authMenus := map[uint]string{}
	for _, menu := range menus {
		if menu.Type == models.MenuMenu {
			var paths []string
			path := "/" + menu.Tag
			paths = menuPath(paths, menus, menu.ParentID)
			for _, p := range paths {
				path =  "/" + p + path
			}
			authMenus[menu.ID] = path
		}
	}

	roles, _ :=  s.repo.GetRolePage(0, -1, map[string]string{})
	for _, role := range roles {
		var mids []uint
		for _, menu := range role.Menu {
			if menu.Type == models.MenuAuth && menu.Tag == models.MethodGet {
				mids = append(mids, menu.ParentID)
			}
		}
		if len(mids) > 0 {
			for _, m := range mids {
				if _, ok := res[authMenus[m]]; !ok {
					res[authMenus[m]] = []string{"admin"}
				}
				b := false
				for _, v := range res[authMenus[m]] {
					if v == role.Tag {
						b = true
						break
					}
				}
				if !b {
					res[authMenus[m]] = append(res[authMenus[m]], role.Tag)
				}
			}
		}
	}
	return



}
