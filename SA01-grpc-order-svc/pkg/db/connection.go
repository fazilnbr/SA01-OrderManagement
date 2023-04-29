package db

import (
	"fmt"

	"github.com/fazilnbr/SA01-OrderManagement/SA01-grpc-order-svc/pkg/config"
	"github.com/fazilnbr/SA01-OrderManagement/SA01-grpc-order-svc/pkg/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDatabase(cfg config.Config) (*gorm.DB, error) {
	// psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", cfg.DBHost, cfg.DBUser, cfg.DBName, cfg.DBPort, cfg.DBPassword)
	// Create the MySQL database URL string
	// url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.)
	// Replace the values below with your MySQL database credentials

	// dsn := "your-username:your-password@tcp(your-host:your-port)/your-database?parseTime=true"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	db, dbErr := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	fmt.Println("database", db)

	db.AutoMigrate(
		domain.Order{},
	)


	return db, dbErr
}
