package models

// InvoiceActivity ...
type InvoiceActivity struct {
	BaseModel
	UserID int32  `json:"userId" gorm:"NOT NULL"`
	Action string `json:"action" gorm:"type:varchar(255);NOT NULL"`

	ErrorMessage  string `json:"errorMessage"`
	SuccessStatus bool   `json:"successStatus"`
}

// TableName ...
func (InvoiceActivity) TableName() string {
	return "invoice_service_activity_logs"
}
