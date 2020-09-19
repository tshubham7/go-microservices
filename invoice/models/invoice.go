package models

import "time"

// BaseModel ...
type BaseModel struct {
	ID        int32      `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt" sql:"index"`
}

// Invoice ...
type Invoice struct {
	BaseModel
	UserID int32  `json:"userId" gorm:"NOT NULL"`
	Action string `json:"action" gorm:"type:varchar(255);NOT NULL"`
}
