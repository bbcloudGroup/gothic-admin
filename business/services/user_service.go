package services

import (
	"errors"
	"gothic-admin/business/models"
	"gothic-admin/business/repositories"
	"gothic-admin/utils"
	"strconv"
	"strings"
)


type UserService interface {
	Login(params models.LoginParams) (models.User, error)

	Add(params models.UserForm) (models.User, error)
	Update(params models.UserForm) (models.User, error)
	Delete(params models.UserForm) error
	ResetPassword(params models.UserForm) (models.User, error)
	StatusChange(params models.UserForm) (models.User, error)
	Get(userID uint) models.User
	GetPage(current int, pageSize int, params map[string]string) models.Page
	CheckMenu(userID uint, target models.Menu) bool
}

func NewUserService(repo repositories.AdminRepo) UserService {
	return &userService{repo:repo}
}

type userService struct {
	repo repositories.AdminRepo
}


func (s *userService) Login(params models.LoginParams) (user models.User, err error) {

	err = nil

	if params.Type == "account" {
		user = s.repo.GetUserByMail(params.UserName)
	} else {
		user = s.repo.GetUserByMobile(params.Mobile)
	}
	if user.ID == 0 {
		err = errors.New("用户不存在")
		return
	}
	if !user.Status {
		err = errors.New("用户已被禁用")
		return
	}

	if params.Type == "account" && !(utils.Md5s(params.Password) == user.Password) {
		err = errors.New("密码错误")
		return
	}

	token, ok := utils.JwtToken(&utils.Claims{
		UserID:         user.ID,
		Name:           user.Name,
		Mail:           user.Mail,
		Avatar:         user.Avatar,
		Mobile:         user.Mobile,
	})

	if ok {
		s.repo.SetToken(&user, token)
		return
	}

	err = errors.New("")
	return
}


func (s *userService) Get(userID uint) models.User {
	return s.repo.GetUserByID(userID)
}


func (s *userService) GetPage(current int, pageSize int, params map[string]string) models.Page {

	var data []interface{}

	users, count := s.repo.GetUserPage(current, pageSize, params)

	for _, user := range users{
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

func (s *userService) Add(params models.UserForm) (user models.User, err error) {
	if !(s.repo.GetUserByMail(params.Mail).ID == 0) {
		err = errors.New("该邮箱已经被注册了")
		return
	}

	if !(s.repo.GetUserByMobile(params.Mobile).ID == 0) {
		err = errors.New("该手机已经被注册了")
		return
	}

	user, ok := s.repo.AddUser(params)
	if !ok {
		err = errors.New("用户添加失败")
	}
	return
}


func (s *userService) Update(params models.UserForm) (user models.User, err error) {
	return s.repo.UpdateUser(params)
}

func (s *userService) ResetPassword(params models.UserForm) (models.User, error) {
	params.Password = "123456"
	updateRule := func(user *models.User, params models.UserForm) {
		user.Password = utils.Md5s(params.Password)
	}
	return s.repo.UpdateUser(params, updateRule)
}


func (s *userService) StatusChange(params models.UserForm) (models.User, error) {
	updateRule := func(user *models.User, params models.UserForm) {
		user.Status = !user.Status
	}
	return s.repo.UpdateUser(params, updateRule)
}

func (s *userService) Delete(params models.UserForm) (err error) {

	if len(params.IDS) == 0 {
		return
	}

	for _, id := range params.IDS {
		if id == 1 {
			err = errors.New("无法删除超级管理员用户")
			return
		}
	}

	return s.repo.DeleteUser(params.IDS)

}


func (s *userService) CheckMenu(userID uint, target models.Menu) bool {

	user := s.repo.GetUserByID(userID)
	var ids []string
	for _, role := range user.Role {
		if role.ID == 1 {
			return true
		}
		ids = append(ids, strconv.Itoa(int(role.ID)))
	}
	roles, _ := s.repo.GetRolePage(0, -1, map[string]string{
		"id": strings.Join(ids, ","),
	})
	for _, role := range roles {
		for _, menu := range role.Menu {
			if menu.ID == target.ID {
				return true
			}
		}
	}
	return false


}
