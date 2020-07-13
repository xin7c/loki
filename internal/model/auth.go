package model

import "github.com/jinzhu/gorm"

type Auth struct {
	*Model
	Username string `json:"username"`
	Password string `json:"password"`
}

type Auths []Auth

func (a Auth) Create(db *gorm.DB) error {
	return db.Create(&a).Error
}

func (as Auths) Find(db *gorm.DB) (Auths, error) {
	err := db.Find(&as).Error
	if err != nil {
		return nil, err
	}
	return as, nil
}

func (a Auth) TableName() string {
	return "loki_auth"
}
