package services

import (
	"github.com/go-playground/validator"
	"github.com/tshubham7/go-microservices/user/models"
	"github.com/tshubham7/go-microservices/user/repository"
)

type user struct {
	a repository.UserRepo
}

// UserUpdateParams ...
type UserUpdateParams struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"email"`
}

// UserCreateRequest ...
type UserCreateRequest struct {
	UserUpdateParams
	Phone string `json:"phone" validate:"required"`
}

// Validate ...
func (u UserCreateRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}

// ToModel ...
func (u UserCreateRequest) ToModel() *models.User {
	// perform any change in the state
	return &models.User{
		Name:  u.Name,
		Email: u.Email,
		Phone: u.Phone,
	}
}

// UserListQueryParams ...
type UserListQueryParams struct {
	Sort   string
	Order  string
	Limit  int32
	Offset int32
}

// UserService ...
type UserService interface {
	// create new user
	Create(Request UserCreateRequest) (*models.User, error)

	// list all users
	ListAll(queries *UserListQueryParams) ([]models.User, error)

	// delete user
	Delete(id int32) error

	// update user
	Update(id int32, req UserUpdateParams) (*models.User, error)
}

// NewUserService ...
func NewUserService(a repository.UserRepo) UserService {
	return &user{a}
}

// Create ...
func (usr user) Create(user UserCreateRequest) (*models.User, error) {

	u := user.ToModel()
	err := usr.a.Create(u).Error

	return u, err
}

// ListAll ...
func (usr user) ListAll(q *UserListQueryParams) ([]models.User, error) {
	return usr.a.ListAll(q.Sort, q.Order, q.Limit, q.Offset)
}

// Delete ...
func (usr user) Delete(id int32) error {
	return usr.a.Delete(id)
}

// Update ...
func (usr user) Update(id int32, req UserUpdateParams) (*models.User, error) {
	return usr.a.Update(id, req.Name, req.Email)
}
