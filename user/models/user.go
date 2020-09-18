package models

import (
	"github.com/jinzhu/gorm"
)

// User ...
type User struct {
	gorm.Model
	Name  string `gorm:"type:varchar(255);NOT NULL"`
	Email string `gorm:"type:varchar(255);NOT NULL;UNIQUE;UNIQUE INDEX"`
	Phone string `gorm:"type:varchar(255);NOT NULL;UNIQUE;UNIQUE INDEX"`
}

// Users ...
type Users []User
