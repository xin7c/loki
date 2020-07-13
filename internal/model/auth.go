package model

import "github.com/jinzhu/gorm"

type Auth struct {
	ID         int `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (a Auth) Create(db *gorm.DB) error{
	return db.Create(&a).Error
}

func (a Auth) TableName() string{
	return "loki_auth"
}