package model

import (
	"database/sql"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string         `json:"username" gorm:"not null"`
	Password string         `json:"password" gorm:"not null"`
	UserType sql.NullString `json:"user_type" gorm:"column:user_type;default:'dev'"`
}

func (user User) Add(db *gorm.DB) error {
	return db.Create(&user).Error
}

func (user User) Modify(db *gorm.DB, encodePWD string) error {
	return db.Model(&user).Update("password", encodePWD).Error
}

func (user User) GetUsers(db *gorm.DB) error {
	return db.Create(&user).Error
}

func (user *User) CheckUserExist(db *gorm.DB) error {
	return db.
		Where("username = ?", user.Username).
		First(&user).
		Error
}

func (user *User) GetUserInfo(db *gorm.DB) error {
	return db.Where("username = ?", user.Username).
		First(&user).
		Error
}

type Users []User

func (users Users) Find(db *gorm.DB) (Users, error) {
	err := db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}
