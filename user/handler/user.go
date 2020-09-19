package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	protos "github.com/tshubham7/go-microservices/invoice/protos/invoice"
	"github.com/tshubham7/go-microservices/user/repository"
	"github.com/tshubham7/go-microservices/user/services"
)

type user struct {
	u       repository.UserRepo
	l       *log.Logger
	invoice *services.Invoice
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
func NewUserHandler(u repository.UserRepo, lr repository.LogRepo, l *log.Logger, cc protos.InvoiceClient) UserHandler {
	ls := services.NewLogService(lr, l)
	in := services.NewInvoiceService(l, cc, ls)
	return &user{u: u, l: l, invoice: in}
}

// Create ...
func (u user) Create() gin.HandlerFunc {
	sr := services.NewUserService(u.u, u.l)

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

		usr, err := sr.Create(params)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "failed to create user",
				"error":   err.Error(),
			})
			return
		}

		// call invoice service
		go u.invoice.Create(int32(usr.ID), "create")

		c.JSON(http.StatusOK, usr)
	}
}

// List ...
func (u user) List() gin.HandlerFunc {
	sr := services.NewUserService(u.u, u.l)

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

		resp := ListResponse{
			Limit: q.Limit, Offset: q.Offset, Count: int32(len(u)), Result: u}

		c.JSON(http.StatusOK, resp)
	}
}

// Delete ...
func (u user) Delete() gin.HandlerFunc {
	sr := services.NewUserService(u.u, u.l)

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

		// call invoice service
		go u.invoice.Create(int32(uid), "delete")

		c.JSON(http.StatusOK, gin.H{"message": "successfully deleted"})
	}
}

// Detail ...
func (u user) Update() gin.HandlerFunc {
	sr := services.NewUserService(u.u, u.l)

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

		// call invoice service
		go u.invoice.Create(int32(uid), "update")

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

// ListResponse ...
type ListResponse struct {
	Limit  int32       `json:"limit"`
	Offset int32       `json:"offset"`
	Count  int32       `json:"count"`
	Result interface{} `json:"results"`
}
