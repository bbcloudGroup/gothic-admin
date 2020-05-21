package models

const (
	MethodGet = "get"
	MethodAdd = "add"
	MethodUpdate = "update"
	MethodDelete = "delete"
)

type Form struct {
	Method string `json:"method"`
	ID		uint  	`json:"id"`
	IDS 	[]uint	`json:"ids"`
}

type UserForm struct {
	Form
	Mail	string 	`json:"mail"`
	Mobile 	string 	`json:"mobile"`
	Password string `json:"password"`
	Name 	string 	`json:"name"`
	Role 	[]Role 	`json:"roles"`
}

type RegisterParams struct {
	UserForm
	Confirm		string 	`json:"confirm"`
	Captcha		string 	`json:"captcha"`
}

type LoginParams struct {
	UserName	string 	`json:"userName"`
	Password 	string 	`json:"password"`
	Mobile		string 	`json:"mobile"`
	Captcha 	string 	`json:"captcha"`
	Type 		string 	`json:"type"`
}

type RoleForm struct {
	Form
	Name 	string 	`json:"name"`
	Tag 	string 	`json:"tag"`
	Menu	[]string	`json:"menus"`
}

type MenuForm struct {
	Form
	Name 	string 	`json:"name"`
	Tag 	string 	`json:"tag"`
	Type	int 	`json:"type"`
	ParentID uint 	`json:"parent_id"`
	Children []MenuForm `json:"children"`
}

