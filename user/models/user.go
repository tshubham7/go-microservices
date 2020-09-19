package models

import "time"

// BaseModel ...
type BaseModel struct {
	ID        int32      `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt" sql:"index"`
}

// User ...
type User struct {
	BaseModel
	Name  string `json:"name" gorm:"type:varchar(255);NOT NULL"`
	Email string `json:"email" gorm:"type:varchar(255);NOT NULL;UNIQUE;UNIQUE INDEX"`
	Phone string `json:"phone" gorm:"type:varchar(255);NOT NULL;UNIQUE;UNIQUE INDEX"`
}

// Users ...
type Users []User
