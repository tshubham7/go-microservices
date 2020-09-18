package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tshubham7/go-microservices/invoice/repository"
	"github.com/tshubham7/go-microservices/invoice/services"
)

type invoice struct {
	r repository.InvoiceRepo
	l *log.Logger
}

// InvoiceHandler ...
type InvoiceHandler interface {
	// create user
	Create() gin.HandlerFunc

	// list user
	List() gin.HandlerFunc

	// delete user
	Delete() gin.HandlerFunc
}

// NewInvoiceHandler ...
func NewInvoiceHandler(r repository.InvoiceRepo, l *log.Logger) InvoiceHandler {
	return &invoice{r: r, l: l}
}

// Create ...
func (in invoice) Create() gin.HandlerFunc {
	sr := services.NewInvoiceService(in.r)

	return func(c *gin.Context) {
		var err error

		var params services.InvoiceCreateRequest
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
				"message": "failed to create invoice",
				"error":   err.Error(),
			})
			return
		}

		// call invoice service
		c.JSON(http.StatusOK, u)
	}
}

// List ...
func (in invoice) List() gin.HandlerFunc {
	sr := services.NewInvoiceService(in.r)

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
				"message": "failed to fetch invoices",
				"error":   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, u)
	}
}

// Delete ...
func (in invoice) Delete() gin.HandlerFunc {
	sr := services.NewInvoiceService(in.r)

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
				"message": "failed to delete invoice",
				"error":   err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "successfully deleted"})
	}
}

// validateQueries validating query params

/*validateQueries
validating query params
return limit, offset, sort, order and err
*/
func validateQueries(args ...string) (*services.InvoiceListQueryParams, error) {
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

	return &services.InvoiceListQueryParams{
		Sort:   sort.String(),
		Order:  order.String(),
		Limit:  limit.Int(),
		Offset: offset.Int()}, nil
}
