package models

import (
	"time"
)



const (
	MenuGroup = 1
	MenuMenu  = 2
	MenuAuth  = 3

)

type User struct {
	ID        	uint `gorm:"primary_key" json:"id"`
	CreatedAt 	time.Time	`json:"created_at"`
	UpdatedAt 	time.Time	`json:"-"`
	DeletedAt 	*time.Time 	`sql:"index" json:"-"`
	Password 	string	`gorm:"size:100;not null;" json:"-"`
	Name 	 	string	`gorm:"size:100;not null;" json:"name"`
	Avatar		string  `gorm:"size:255" json:"avatar"`
	Mail 		string 	`gorm:"size:50;index" json:"mail"`
	Mobile		string 	`gorm:"size:11;index" json:"mobile"`
	Role 		[]Role 	`gorm:"many2many:user_role;" json:"roles"`
	Status		bool	`json:"status"`
	RememberToken string `gorm:"size:100"`
}

type Role struct {
	ID			uint `gorm:"primary_key" json:"id"`
	CreatedAt 	time.Time	`json:"-"`
	UpdatedAt 	time.Time	`json:"-"`
	DeletedAt 	*time.Time 	`sql:"index" json:"-"`
	Name 		string 	`gorm:"size:100;not null;" json:"name"`
	Tag 		string 	`gorm:"size:100;not null;" json:"tag"`
	Menu		[]Menu 	`gorm:"many2many:role_menu;" json:"menus"`
}

type Menu struct {
	ID			uint 		`gorm:"primary_key" json:"id"`
	CreatedAt 	time.Time	`json:"-"`
	UpdatedAt 	time.Time	`json:"-"`
	DeletedAt 	*time.Time 	`sql:"index" json:"-"`
	Tag			string		`gorm:"size:255" json:"tag"`	// 标识
	Name 		string 		`gorm:"size:255" json:"name"`	// 名称
	Type		int			`json:"type"`					// 类型：组、菜单、权限
	ParentID	uint		`json:"parent_id"`
	Children 	[]Menu		`gorm:"-" json:"children"`
}

type Log struct {
	ID			uint 		`gorm:"primary_key" json:"id"`
	CreatedAt 	time.Time	`json:"created_at"`
	UpdatedAt 	time.Time	`json:"-"`
	DeletedAt 	*time.Time 	`sql:"index" json:"-"`
	User		User		`json:"user"`
	UserID 		uint		`json:"-"`
	Menu		Menu		`json:"menu"`
	MenuID		uint		`json:"-"`
	Method 		Menu 		`json:"method"`
	MethodID 	uint		`json:"-"`
	Data 		string		`gorm:"size:65535" json:"data"`
}