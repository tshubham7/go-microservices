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
	// list user
	List() gin.HandlerFunc
}

// NewInvoiceHandler ...
func NewInvoiceHandler(r repository.InvoiceRepo, l *log.Logger) InvoiceHandler {
	return &invoice{r: r, l: l}
}

// List ...
func (in invoice) List() gin.HandlerFunc {

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

		uid := c.Query("userID")
		if uid == "" {
			// list all invoices
			in.listAll(c, q)
			return
		}

		// else list invoice by user
		id, err := strconv.Atoi(uid)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "missing or invalid params: userID",
				"error":   err.Error(),
			})
			return
		}

		in.listByUser(c, q, int32(id))
	}
}

// listAll ...
func (in invoice) listAll(c *gin.Context, queries *services.InvoiceListQueryParams) {
	sr := services.NewInvoiceService(in.r)
	u, err := sr.ListAll(queries)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to fetch invoices",
			"error":   err.Error(),
		})
		return
	}
	resp := ListResponse{
		Limit: queries.Limit, Offset: queries.Offset, Count: int32(len(u)), Result: u}

	c.JSON(http.StatusOK, resp)
}

// listByUser ...
func (in invoice) listByUser(c *gin.Context, queries *services.InvoiceListQueryParams, uid int32) {
	sr := services.NewInvoiceService(in.r)
	u, err := sr.ListByUser(queries, uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to fetch invoices",
			"error":   err.Error(),
		})
		return
	}
	resp := ListResponse{
		Limit: queries.Limit, Offset: queries.Offset, Count: int32(len(u)), Result: u}

	c.JSON(http.StatusOK, resp)
}

// ListResponse ...
type ListResponse struct {
	Limit  int32       `json:"limit"`
	Offset int32       `json:"offset"`
	Count  int32       `json:"count"`
	Result interface{} `json:"results"`
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
