package services

import (
	"log"

	"github.com/tshubham7/go-microservices/user/models"
	"github.com/tshubham7/go-microservices/user/repository"
)

type logs struct {
	lg repository.LogRepo
	l  *log.Logger
}

// LogService ...
type LogService interface {
	// create new service log
	Create(request models.InvoiceActivity) (*models.InvoiceActivity, error)

	// list all logs
	ListAll(limit int32, offset int32) ([]models.InvoiceActivity, error)
}

// NewLogService ...
func NewLogService(r repository.LogRepo, l *log.Logger) LogService {
	return &logs{r, l}
}

// Create ...
func (lg logs) Create(req models.InvoiceActivity) (*models.InvoiceActivity, error) {

	err := lg.lg.Create(&req).Error

	return &req, err
}

// ListAll ...
func (lg logs) ListAll(lmt, ofs int32) ([]models.InvoiceActivity, error) {
	return lg.lg.ListAll(lmt, ofs)
}
