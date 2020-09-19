package repository

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	"github.com/tshubham7/go-microservices/invoice/models"
)

type invoice struct {
	db *gorm.DB
	l  *log.Logger
}

// InvoiceRepo ..
type InvoiceRepo interface {
	// create new invoice
	Create(user *models.Invoice) *gorm.DB

	// list all invoice
	ListAll(sort, order string, limit, offset int32) ([]models.Invoice, error)

	// list invoice by user id
	ListByUser(sort, order string, limit, offset, userID int32) ([]models.Invoice, error)
}

// NewInvoiceRepo ...
func NewInvoiceRepo(db *gorm.DB, l *log.Logger) InvoiceRepo {
	return &invoice{db, l}
}

// Create ...
func (in *invoice) Create(invoice *models.Invoice) *gorm.DB {
	return in.db.Create(invoice)
}

// ListAll ...
func (in *invoice) ListAll(sort, order string, limit, offset int32) ([]models.Invoice, error) {
	var invoices = []models.Invoice{}

	result := in.db.Table("invoices").Order(fmt.Sprintf("%s %s", sort, order))
	result = result.Limit(limit).Offset(offset).Find(&invoices)

	return invoices, result.Error
}

// ListByUser ...
func (in *invoice) ListByUser(sort, order string, limit, offset, uid int32) ([]models.Invoice, error) {
	var invoices = []models.Invoice{}

	result := in.db.Table("invoices").Where("user_id = ?", uid).Order(fmt.Sprintf("%s %s", sort, order))
	result = result.Limit(limit).Offset(offset).Find(&invoices)

	return invoices, result.Error
}
