package admin

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"gothic-admin/business/models"
	"gothic-admin/business/services"
	"gothic-admin/utils"
	"gothic-admin/web/middleware"
	"net/http"
	"strconv"
	"time"
)

type ApiController struct{
	User services.UserService
	Captcha services.CaptchaService
	Menu services.MenuService
	Role services.RoleService
}

func NewApiController(
	captcha services.CaptchaService,
	user services.UserService,
	menu services.MenuService,
	role services.RoleService) ApiController {
	return ApiController{Captcha:captcha, User:user, Menu:menu, Role:role}
}


func (c *ApiController) BeforeActivation(b mvc.BeforeActivation){
	b.Handle("GET", "/captcha/{mobile:string}","GetCaptcha")
	b.Handle("GET", "/current", "GetCurrent", middleware.Jwt)
	b.Handle("GET", "/menu", "GetMenu", middleware.Jwt)
	b.Handle("GET", "/account", "GetAccount", middleware.Jwt)
}

func (c *ApiController) GetMenu() (res struct {
	models.Response
	Data 	map[string][]string	`json:"data"`
}){
	res.Status = "ok"
	res.Message = ""
	res.Data = c.Role.GetMenuRoles()
	return
}

func (c *ApiController) GetMenus(ctx iris.Context) models.Page {
	return c.Menu.GetPage(0, -1, ctx.URLParams())
}

func (c *ApiController) GetRoles(ctx iris.Context) models.Page {
	return c.Role.GetPage(0, -1, ctx.URLParams())
}

func (c *ApiController) GetUsers(ctx iris.Context) models.Page {
	return c.User.GetPage(0, -1, ctx.URLParams())
}

func (c *ApiController) GetCurrent(ctx iris.Context) models.User {
	return c.User.Get(getUserID(ctx))
}

func (c *ApiController) GetAccount(ctx iris.Context) models.AccountInfo {
	user := c.User.Get(getUserID(ctx))
	return models.AccountInfo{
		Mail:		 user.Mail,
		Name:        user.Name,
		Avatar:      user.Avatar,
		ID:          strconv.Itoa(int(user.ID)),
		Signature:   "海纳百川，有容乃大",
		Title:       "交互专家",
		Group:       "蚂蚁金服－某某某事业群－某某平台部－某某技术部－UED",
		Tag:         nil,
		NotifyCount: 12,
		UnreadCount: 11,
		Country:     "China",
		Geographic: struct {
			Province models.Tag `json:"province"`
			City     models.Tag `json:"city"`
		}{
			Province: models.Tag{
				Label:	"浙江省",
				Key: 	"330000",
			},
			City:	  models.Tag{
				Label:	"杭州市",
				Key: 	"330100",
			},
		},
		Address: "西湖区工专路 77 号",
		Phone:   "0752-268888888",
	}
}


func (c *ApiController) GetCaptcha(mobile string) (res models.Response) {

	if !c.Captcha.IsPhoneNumber(mobile) {
		res.Status = "error"
		res.Message = "手机号错误"
		return
	}

	captcha := c.Captcha.SendCaptcha(mobile)
	res.Status = "ok"
	res.Message = "验证码已发出" + captcha
	return
}

func (c *ApiController) PostRegister(ctx iris.Context) (res struct{
	models.Response
	CurrentAuthority	string 	`json:"currentAuthority"`
}) {

	var params models.RegisterParams
	err := ctx.ReadJSON(&params)
	if err != nil {
		panic(err)
	}

	res.CurrentAuthority = "guest"

	if !c.Captcha.Check(params.Mobile, params.Captcha) {
		res.Status = "error"
		res.Message = "验证码错误"
		return
	}

	_, err = c.User.Add(params.UserForm)
	if err != nil {
		res.Status = "error"
		res.Message = err.Error()
		return
	}

	res.Status = "ok"
	res.Message = "注册成功"
	return
}

func (c *ApiController) PostLogin(ctx iris.Context) (res struct{
	models.Response
	Type 	string 	`json:"type"`
	CurrentAuthority	string 	`json:"currentAuthority"`
	Token 	string 	`json:"token"`
	UserID 	int 	`json:"uid"`
}) {

	var params models.LoginParams
	err := ctx.ReadJSON(&params)
	if err != nil {
		panic(err)
	}

	res.Type = params.Type
	res.CurrentAuthority = "guest"

	if params.Type != "account" {
		if !c.Captcha.Check(params.Mobile, params.Captcha) {
			res.Status = "error"
			res.Message = "验证码错误"
			return
		}
	}

	user, err := c.User.Login(params)
	if err != nil {
		res.Status = "error"
		res.Message = err.Error()
		return
	}


	var role string
	for _, r:= range user.Role {
		if len(role) == 0 {
			role = "\"" + r.Tag + "\""
		} else {
			role = role + ",\"" + r.Tag + "\""
		}
	}

	// 登录时间
	lt := strconv.FormatInt(time.Now().UnixNano(),10)
	// sid
	sid := utils.Md5s("Bearer " + user.RememberToken + utils.Md5s("[" + role + "]" + strconv.Itoa(int(user.ID))) + lt)
	ctx.SetCookieKV("_lt", lt, func(cookie *http.Cookie) {
		cookie.HttpOnly = false
	})
	ctx.SetCookieKV("_sid", sid, func(cookie *http.Cookie) {
		cookie.HttpOnly = false
	})

	res.Token = user.RememberToken
	res.Status = "ok"
	res.Message = "登录成功"
	res.CurrentAuthority = "[" + role + "]"
	res.UserID = int(user.ID)
	return

}
