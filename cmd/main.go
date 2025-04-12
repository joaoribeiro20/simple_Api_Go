package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joaoribeiro20/simplegoapi/config"
	"github.com/joaoribeiro20/simplegoapi/models"
	"github.com/joaoribeiro20/simplegoapi/routes"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	config.ConnectDatabase()
	config.DB.AutoMigrate(&models.Budget{})

	r := gin.Default()
	routes.RegisterBudgetRoutes(r)

	r.Run(":8080")
}
