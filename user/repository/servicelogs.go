package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/tshubham7/go-microservices/user/models"
)

type logs struct {
	db *gorm.DB
}

// LogRepo ..
type LogRepo interface {
	// create new service log
	Create(logRequest *models.InvoiceActivity) *gorm.DB

	// list all service logs
	ListAll(limit, offset int32) ([]models.InvoiceActivity, error)
}

// NewLogRepo ...
func NewLogRepo(db *gorm.DB) LogRepo {
	return &logs{db}
}

// Create ...
func (lg *logs) Create(req *models.InvoiceActivity) *gorm.DB {
	return lg.db.Create(req)
}

// ListAll ...
func (lg *logs) ListAll(limit, offset int32) ([]models.InvoiceActivity, error) {
	var lgs = []models.InvoiceActivity{}

	result := lg.db.Table("invoice_service_activity_logs").Limit(limit).Offset(offset).Find(&lgs)

	return lgs, result.Error
}
