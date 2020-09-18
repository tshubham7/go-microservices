package services

import (
	"github.com/go-playground/validator"
	"github.com/tshubham7/go-microservices/invoice/models"
	"github.com/tshubham7/go-microservices/invoice/repository"
)

type invoice struct {
	r repository.InvoiceRepo
}

// InvoiceCreateRequest ...
type InvoiceCreateRequest struct {
	UserID int32  `json:"user_id" validate:"required"`
	Action string `json:"action" validate:"action"`
}

// Validate ...
func (req InvoiceCreateRequest) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("action", ValidateAction)
	return validate.Struct(req)
}

var actionMap = map[string]bool{
	"update": true,
	"create": true,
	"delete": true,
}

// ValidateAction ...
func ValidateAction(fl validator.FieldLevel) bool {
	if actionMap[fl.Field().String()] {
		return true
	}
	return false
}

// ToModel ...
func (req InvoiceCreateRequest) ToModel() *models.Invoice {
	// perform any change in the state
	return &models.Invoice{
		UserID: req.UserID,
		Action: req.Action,
	}
}

// InvoiceListQueryParams ...
type InvoiceListQueryParams struct {
	Sort   string
	Order  string
	Limit  int32
	Offset int32
}

// InvoiceService ...
type InvoiceService interface {
	// create new invoice
	Create(Request InvoiceCreateRequest) (*models.Invoice, error)

	// list all invoices
	ListAll(queries *InvoiceListQueryParams) ([]models.Invoice, error)

	// delete invoice
	Delete(id int32) error
}

// NewInvoiceService ...
func NewInvoiceService(a repository.InvoiceRepo) InvoiceService {
	return &invoice{a}
}

// Create ...
func (in invoice) Create(invoice InvoiceCreateRequest) (*models.Invoice, error) {

	u := invoice.ToModel()
	err := in.r.Create(u).Error

	return u, err
}

// ListAll ...
func (in invoice) ListAll(q *InvoiceListQueryParams) ([]models.Invoice, error) {
	return in.r.ListAll(q.Sort, q.Order, q.Limit, q.Offset)
}

// Delete ...
func (in invoice) Delete(id int32) error {
	return in.r.Delete(id)
}
