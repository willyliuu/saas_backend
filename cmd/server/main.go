package main

import (
	"log"

	"gin/internal/config"
	"gin/internal/database"
	"gin/internal/migrations"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()

	database.ConnectDatabase(cfg.DatabaseURL)

	m := migrations.Migrations()
	if err := m.Migrate(); err != nil {
		log.Fatalf("Could not migrate: %v", err)
	}

	log.Println("✅ Database migrated successfully")

	// Setup Gin
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	log.Printf("🚀 Server running on port %s", cfg.PORT)
	r.Run(":" + cfg.PORT)
}
