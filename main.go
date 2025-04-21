package main

import (
	"authsvc/config"
	"authsvc/db"
	"authsvc/handlers"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Ошибка при загрузке конфига: %v", err)
	}
	dbConn, err := db.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Ошибка подключения к БД: %v", err)
	}
	handlers.DB = dbConn
	handlers.JWTSecret = cfg.JWTSecret
	r := gin.Default()
	r.POST("/token", handlers.Issue)
	r.POST("/refresh", handlers.Refresh)
	log.Println("Сервис запущен на :8080")
	r.Run(":8080")
}
