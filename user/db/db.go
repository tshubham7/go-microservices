package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // Only for init
	"github.com/tshubham7/go-microservices/user/models"
)

type dbConfig struct {
	host     string
	port     int
	user     string
	dbname   string
	password string
}

// var config = dbConfig{"localhost", 5432, "root", "user_service_db", "1234"}

// func getDatabaseURL() string {
// 	return fmt.Sprintf(
// 		"host=%s port=%d user=%s dbname=%s password=%s",
// 		config.host, config.port, config.user, config.dbname, config.password)
// }

// GetDatabase ...
func GetDatabase() (*gorm.DB, error) {
	// db, err := gorm.Open("mysql", getDatabaseURL())
	db, err := gorm.Open("sqlite3", "user_service.db")
	return db, err
}

// RunMigrations ...
// running migrations
func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.InvoiceActivity{})
}
