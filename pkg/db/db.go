package db

import (
	"lendral3n/KulinerKlasik-Microservices-gRPC-Auth/pkg/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func Init(url string) Handler {
	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.User{})
	return Handler{db}
}