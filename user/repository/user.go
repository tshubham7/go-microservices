package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/tshubham7/go-microservices/user/models"
)

type user struct {
	db *gorm.DB
}

// UserRepo ..
type UserRepo interface {
	// create new user
	Create(user *models.User) *gorm.DB

	// list all users
	ListAll(sort, order string, limit, offset int32) ([]models.User, error)

	// update user
	Update(id int32, name string, email string) (*models.User, error)

	// delete user
	Delete(id int32) error
}

// NewUserRepo ...
func NewUserRepo(db *gorm.DB) UserRepo {
	return &user{db}
}

// Create ...
func (u *user) Create(user *models.User) *gorm.DB {
	return u.db.Create(user)
}

// Delete ...
func (u *user) Delete(id int32) error {

	result := u.db.Table("users").Where("id=?", id).Delete(&models.User{})

	return result.Error
}

// Update ...
func (u *user) Update(id int32, name, email string) (*models.User, error) {
	var user models.User
	u.db.Where("id = ?", id).Find(&user)

	if name != "" {
		user.Name = name
	}

	if email != "" {
		user.Email = email
	}

	err := u.db.Save(user).Error
	return &user, err
}

// ListAll ...
func (u *user) ListAll(sort, order string, limit, offset int32) ([]models.User, error) {
	var users = []models.User{}

	result := u.db.Table("users").Order(fmt.Sprintf("%s %s", sort, order))
	result = result.Limit(limit).Offset(offset).Find(&users)

	return users, result.Error
}
