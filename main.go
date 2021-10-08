package main

import (
	"os"

	"TwitterClone/api/db"
	"TwitterClone/api/controllers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	app := gin.Default()
	db := db.PostgresDB{}
	mountController := controllers.MountController{}

	db.InitializeDatabase()
	mountController.Init(app.Group("/api/v1"), db.DB)

	app.Run(":" + os.Getenv("PORT"))
}
