package services

import (
	"context"
	"fmt"
	"log"

	protos "github.com/tshubham7/go-microservices/invoice/protos/invoice"
)

// Invoice ...
type Invoice struct {
	l  *log.Logger
	cc protos.InvoiceClient
}

// NewInvoiceService ...
func NewInvoiceService(l *log.Logger, cc protos.InvoiceClient) *Invoice {
	return &Invoice{l, cc}
}

// Create ...
// create invoice through grpc
func (in Invoice) Create(userID int32, action string) {
	rr := protos.CreateRequest{
		UserID: userID,
		Action: action,
	}

	resp, err := in.cc.Create(context.Background(), &rr)
	if err != nil {
		// here you use your mechanism to keep track of errors
		// what i could here is entry these errors in the database
		in.l.Println(err)
	}
	fmt.Println(resp.Message, resp.Status)
}
