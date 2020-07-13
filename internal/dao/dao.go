package dao

import "github.com/jinzhu/gorm"

type Dao struct {
	engine *gorm.DB
}