package services

import (
	"errors"
	"gothic-admin/business/models"
	"gothic-admin/business/repositories"
)

type MenuService interface {
	Add(params models.MenuForm) (models.Menu, error)
	Update(params models.MenuForm) (models.Menu, error)
	Delete(params models.MenuForm) error
	GetPage(current int, pageSize int, params map[string]string) models.Page
	FindMenu(tags []string, method string) (models.Menu, models.Menu)
}

func NewMenuService(repo repositories.AdminRepo) MenuService {
	return &menuService{repo:repo}
}

type menuService struct {
	repo repositories.AdminRepo
}


func (s *menuService) Add(params models.MenuForm) (menu models.Menu, err error) {

	if s.repo.GetMenuBy(params).ID != 0 {
		err = errors.New("该菜单已经存在，不要重复添加")
		return
	}

	parent := s.repo.GetMenuByID(params.ParentID)

	if params.Type == models.MenuAuth && parent.Type != models.MenuMenu {
		err = errors.New("只能在菜单下添加权限，选择一个菜单")
		return
	}

	if params.Type == models.MenuMenu && parent.Type != models.MenuGroup {
		err = errors.New("只能在组下添加菜单，请选择一个分组")
		return
	}

	if params.Type == models.MenuGroup && parent.Type != models.MenuGroup && parent.Type != 0 {
		err = errors.New("只能在组下添加分组，请选择一个分组")
		return
	}

	menu, ok := s.repo.AddMenu(params)
	if !ok {
		err = errors.New("菜单添加失败")
	}
	return
}

func (s *menuService) Update(params models.MenuForm) (menu models.Menu, err error) {

	if id := s.repo.GetMenuBy(params).ID; id != 0 && id != params.ID{
		err = errors.New("该菜单已经存在，不要重复添加")
		return
	}

	if params.ID == params.ParentID {
		err = errors.New("不能将自己添加到自己")
		return
	}

	parent := s.repo.GetMenuByID(params.ParentID)

	if params.Type == models.MenuAuth && parent.Type != models.MenuMenu {
		err = errors.New("只能在菜单下添加权限，选择一个菜单")
		return
	}

	if params.Type == models.MenuMenu && parent.Type != models.MenuGroup {
		err = errors.New("只能在组下添加菜单，请选择一个组")
		return
	}

	if params.Type == models.MenuGroup && parent.Type != models.MenuGroup && parent.Type != 0 {
		err = errors.New("只能在组下添加组，请选择一个组")
		return
	}
	return s.repo.UpdateMenu(params)
}

func (s *menuService) Delete(params models.MenuForm) (err error) {
	if len(params.IDS) == 0 {
		return
	}
	return s.repo.DeleteMenu(params.IDS)
}


func toTree(menus []models.Menu, parentId uint) (nodes []models.Menu) {
	for _, menu := range menus {
		if parentId == menu.ParentID {
			node := models.Menu{
				ID:        menu.ID,
				Tag:       menu.Tag,
				Name:      menu.Name,
				Type:      menu.Type,
				ParentID:  menu.ParentID,
				Children:  toTree(menus, menu.ID),
			}
			nodes = append(nodes, node)
		}
	}
	return
}


func (s *menuService) GetPage(current int, pageSize int, params map[string]string) models.Page {
	var data []interface{}

	menus, count :=  s.repo.GetMenuPage(current, pageSize, params)

	for _, user := range toTree(menus, 0) {
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


func (s *menuService) FindMenu(tags []string, method string) (menu models.Menu, target models.Menu) {

	if len(tags) == 0 {
		menu = s.repo.GetMenuBy(models.MenuForm{
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
			menu = s.repo.GetMenuBy(models.MenuForm{
				Tag:      tag,
				Type:     t,
				ParentID: pid,
			})
		}
	}

	if menu.ID > 0 {
		target = s.repo.GetMenuBy(models.MenuForm{
			Tag:      method,
			Type:     models.MenuAuth,
			ParentID: menu.ID,
		})
	}
	return
}