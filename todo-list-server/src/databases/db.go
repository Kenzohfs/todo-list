package databases

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	connectionString := "host=localhost user=admin password=0712 dbname=todolist port=5432 sslmode=disable TimeZone=America/Sao_Paulo"

	db, err := gorm.Open(postgres.Open(connectionString))
	if err != nil {
		log.Panic("error connecting to database")
	}

	return db
}
