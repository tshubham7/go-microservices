package services

import (
	"context"
	"log"

	protos "github.com/tshubham7/go-microservices/invoice/protos/invoice"
	"github.com/tshubham7/go-microservices/user/models"
)

// Invoice ...
type Invoice struct {
	l  *log.Logger
	cc protos.InvoiceClient
	sl LogService
}

// NewInvoiceService ...
func NewInvoiceService(l *log.Logger, cc protos.InvoiceClient, sl LogService) *Invoice {
	return &Invoice{l, cc, sl}
}

// Create ...
// create invoice through grpc
func (in Invoice) Create(userID int32, action string) {
	// service log request param
	invLg := models.InvoiceActivity{UserID: userID, Action: action}

	rr := protos.CreateRequest{
		UserID: userID,
		Action: action,
	}

	resp, err := in.cc.Create(context.Background(), &rr)
	if err != nil {
		invLg.ErrorMessage = err.Error()
		in.l.Println(err)
	}

	if !resp.Status {
		// here you use your mechanism to keep track of errors
		// what i could here is entry these errors in the database
		invLg.ErrorMessage = resp.Message
		in.l.Println(resp.Message)
	}

	invLg.SuccessStatus = resp.Status

	// create service log
	_, err = in.sl.Create(invLg)
	if err != nil {
		in.l.Println(err)
		in.l.Println("error while creating invoice service log")
	}
}
