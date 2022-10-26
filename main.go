package main

import (
	"github.com/carlosm27/jwtGinApi/handlers"
	"github.com/carlosm27/jwtGinApi/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
	port := os.Getenv("PORT")

	db := DbInit()

	models.RepoInterface(db)

	r := gin.Default()

	group := r.Group("/api")

	group.POST("/register", handlers.Register)

	if r.Run(":"+port) != nil {
		log.Printf("Error running at port: %s", port)
	}
}

func DbInit() *gorm.DB {
	db, err := models.Setup()
	if err != nil {
		log.Println("Problem setting up database")
	}
	return db
}