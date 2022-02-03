package model

import (
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type UccUsers struct {
	gorm.Model
	Name     string `gorm:"not null; size:70" json:"name"`
	Email    string `gorm:"not null; size:50" json:"email"`
	Gender   string `gorm:"not null; size:1; check:gender==M || gender==F" json:"gender"`
	Username string `gorm:"unique; not null; size:20" json:"username"`
	Password string `gorm:"not null; size:20" json:"password"`
}

// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&UccUsers{})
	return db
}