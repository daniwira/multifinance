package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitializeDatabase(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// err = db.AutoMigrate(
	// 	&customer.Customer{},
	// 	&limit.Limit{},
	// 	&transaction.Transaction{},
	// )
	// if err != nil {
	// 	return nil, err
	// }

	return db, nil
}
