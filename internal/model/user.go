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

type Users []User

func (user User) Add(db *gorm.DB) error {
	return db.Create(&user).Error
}

func (user User) GetUsers(db *gorm.DB) error {
	return db.Create(&user).Error
}

func (users Users) Find(db *gorm.DB) (Users, error) {
	err := db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}
