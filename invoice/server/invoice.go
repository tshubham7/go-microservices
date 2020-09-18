package server

import (
	"context"

	"github.com/jinzhu/gorm"
	protos "github.com/tshubham7/go-microservices/invoice/protos/invoice"
	"github.com/tshubham7/go-microservices/invoice/repository"
	"github.com/tshubham7/go-microservices/invoice/services"
)

// Invoice ...
type Invoice struct {
	s services.InvoiceService
}

// NewInvoice ...
func NewInvoice(db *gorm.DB) *Invoice {
	ir := repository.NewInvoiceRepo(db)
	return &Invoice{services.NewInvoiceService(ir)}
}

// Create ...
func (c Invoice) Create(ctx context.Context, in *protos.CreateRequest) (*protos.CreateResponse, error) {

	req := services.InvoiceCreateRequest{Action: in.Action, UserID: in.UserID}
	_, err := c.s.Create(req)
	if err != nil {
		return &protos.CreateResponse{
			Message: err.Error(),
			Status:  false,
		}, err
	}

	return &protos.CreateResponse{
		Message: "successfully created",
		Status:  true,
	}, nil
}
