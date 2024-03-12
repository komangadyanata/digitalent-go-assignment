package configs

import (
	"fmt"
	"log"

	"go-assignment2/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	dbHost     = "localhost"
	dbUser     = "root"
	dbPassword = "password"
	dbPort     = "3306"
	dbName     = "orders_by"
	dbCharset  = "utf8mb4"
	db         *gorm.DB
	err        error
)

func StartDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName, dbCharset)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic("failed to connect database :", err)
	}

	db.AutoMigrate(&models.Order{})
	db.AutoMigrate(&models.Item{})
}

func GetDB() *gorm.DB {
	return db
}
