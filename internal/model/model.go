package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"loki/pkg/setting"
)

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}

func NewDBEngine(databaseSettingS *setting.DatabaseSettingS) (*gorm.DB, error) {
	db, err := gorm.Open(databaseSettingS.DBType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		databaseSettingS.Username,
		databaseSettingS.Password,
		databaseSettingS.Host,
		databaseSettingS.DBName,
		databaseSettingS.Charset,
		databaseSettingS.Parsetime))
	if err != nil {
		return nil, err
	}
	db.SingularTable(true)
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	return db, nil
}
