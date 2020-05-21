package datasource

import (
	"github.com/bbcloudGroup/gothic/database"
	"github.com/jinzhu/gorm"
)

type Admin struct {
	*gorm.DB
}

func NewAdmin() Admin {
	return Admin{database.G()}
}

