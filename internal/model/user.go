package model

import (
	"database/sql"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string         `json:"username" gorm:"not null"`
	Password string         `json:"password" gorm:"not null"`
	Usertype sql.NullString `json:"user_type" gorm:"column:user_type;default:'dev'"`
}

func (user User) Add(db *gorm.DB) error {
	return db.Create(&user).Error
}
