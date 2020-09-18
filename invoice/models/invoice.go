package models

import (
	"github.com/jinzhu/gorm"
)

// Invoice ...
type Invoice struct {
	gorm.Model
	UserID int32  `gorm:"NOT NULL"`
	Action string `gorm:"type:varchar(255);NOT NULL"`
}
