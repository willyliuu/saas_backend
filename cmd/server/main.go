package main

import (
	"log"

	"gin/internal/config"
	"gin/internal/database"
	"gin/internal/domain/organization"
	"gin/internal/domain/user"
	"gin/internal/middleware"
	"gin/internal/migrations"
	"gin/internal/routes"
	"gin/internal/token"

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

	// init dependecies
	userRepo := &user.UserRepository{}
	token := &token.TokenMaker{JWTSecret: cfg.JWTSecret, RefreshSecret: cfg.RefreshSecret}
	userService := &user.UserService{Repo: userRepo, Token: token}
	userHandler := &user.UserHandler{Service: userService}

	orgRepo := &organization.OrganizationRepository{}
	orgService := &organization.OrganizationService{Repo: orgRepo}
	orgHandler := &organization.OrganizationHandler{Service: orgService}

	auth := middleware.AuthMiddleware(token)

	r.POST("/register", userHandler.Register)
	r.POST("/login", userHandler.Login)
	r.POST("/refresh", userHandler.Refresh)

	// protected routes
	routes.UserRoutes(r, userHandler, auth)
	routes.OrganizationRoutes(r, orgHandler, auth)

	log.Printf("🚀 Server running on port %s", cfg.PORT)
	r.Run(":" + cfg.PORT)
}
