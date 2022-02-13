package commons

import (
	"github.com/orutrax/go-crud-postgres-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func GetConnection() *gorm.DB {
	dsn := "host=localhost user=postgres password=mysecretpassword dbname=postgres port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	return db
}

func Migrate() {
	db := GetConnection()

	log.Println("Migrating...")

	db.AutoMigrate(&models.Movement{})
}
