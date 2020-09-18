package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tshubham7/go-microservices/user/repository"
	"github.com/tshubham7/go-microservices/user/services"
)

type user struct {
	u repository.UserRepo
	l *log.Logger
}

// UserHandler ...
type UserHandler interface {
	// create user
	Create() gin.HandlerFunc

	// list user
	List() gin.HandlerFunc

	// delete user
	Delete() gin.HandlerFunc

	// update user
	Update() gin.HandlerFunc
}

// NewUserHandler ...
func NewUserHandler(u repository.UserRepo, l *log.Logger) UserHandler {
	return &user{u: u, l: l}
}

// Create ...
func (u user) Create() gin.HandlerFunc {
	sr := services.NewUserService(u.u)

	return func(c *gin.Context) {
		var err error

		var params services.UserCreateRequest
		err = c.Bind(&params)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   err.Error(),
				"message": "missing or invalid params",
			})
			return
		}

		err = params.Validate()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   err.Error(),
				"message": "missing or invalid params",
			})
			return
		}

		u, err := sr.Create(params)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "failed to create user",
				"error":   err.Error(),
			})
			return
		}

		// call invoice service
		c.JSON(http.StatusOK, u)
	}
}

// List ...
func (u user) List() gin.HandlerFunc {
	sr := services.NewUserService(u.u)

	return func(c *gin.Context) {
		q, err := validateQueries(
			c.Query("limit"),
			c.Query("offset"),
			c.Query("sort"),
			c.Query("order"),
		)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "missing or invalid params",
				"error":   err.Error(),
			})
			return
		}

		u, err := sr.ListAll(q)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "failed to fetch users",
				"error":   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, u)
	}
}

// Delete ...
func (u user) Delete() gin.HandlerFunc {
	sr := services.NewUserService(u.u)

	return func(c *gin.Context) {
		id := c.Param("id")

		uid, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "missing or invalid params: id ",
				"error":   err,
			})
			return
		}

		err = sr.Delete(int32(uid))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "failed to delete user",
				"error":   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "successfully deleted"})
	}
}

// Detail ...
func (u user) Update() gin.HandlerFunc {
	sr := services.NewUserService(u.u)

	return func(c *gin.Context) {
		id := c.Param("id")
		uid, err := strconv.Atoi(id)
		if id == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "missing or invalid params: id ",
				"error":   err,
			})
			return
		}

		var params services.UserUpdateParams
		err = c.Bind(&params)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   err.Error(),
				"message": "missing or invalid params",
			})
			return
		}

		a, err := sr.Update(int32(uid), params)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "failed to update user",
				"error":   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, a)
	}
}

// validateQueries validating query params

/*validateQueries
validating query params
return limit, offset, sort, order and err
*/
func validateQueries(args ...string) (*services.UserListQueryParams, error) {
	limit := Limit(args[0])
	if err := limit.Valid(); err != nil {
		return nil, err
	}

	offset := Offset(args[1])
	if err := offset.Valid(); err != nil {
		return nil, err
	}

	sort := Sort(args[2])
	if err := sort.Valid(); err != nil {
		return nil, err
	}

	order := Order(args[3])
	if err := order.Valid(); err != nil {
		return nil, err
	}

	return &services.UserListQueryParams{
		Sort:   sort.String(),
		Order:  order.String(),
		Limit:  limit.Int(),
		Offset: offset.Int()}, nil
}
