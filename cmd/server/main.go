package main

import (
	"log"

	"gin/internal/config"
	"gin/internal/database"
	"gin/internal/domain/invite"
	"gin/internal/domain/organization"
	"gin/internal/domain/project"
	"gin/internal/domain/task"
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

	auth := middleware.AuthMiddleware(token)

	orgRepo := &organization.OrganizationRepository{}
	orgService := &organization.OrganizationService{Repo: orgRepo}
	orgHandler := &organization.OrganizationHandler{Service: orgService}

	inviteRepo := &invite.InviteRepository{}
	inviteService := &invite.InviteService{Repo: inviteRepo}
	inviteHandler := &invite.InviteHandler{Service: inviteService}

	projectRepo := &project.ProjectRepository{}
	projectService := &project.ProjectService{Repo: projectRepo}
	projectHandler := &project.ProjectHandler{Service: projectService}

	taskRepo := &task.TaskRepository{}
	taskService := &task.TaskService{Repo: taskRepo}
	taskHandler := &task.TaskHandler{Service: taskService}

	r.POST("/register", userHandler.Register)
	r.POST("/login", userHandler.Login)
	r.POST("/refresh", userHandler.Refresh)

	// protected routes
	routes.UserRoutes(r, userHandler, auth)
	routes.OrganizationRoutes(r, orgHandler, inviteHandler, projectHandler, auth)
	routes.InviteRoutes(r, inviteHandler, auth)
	routes.ProjectRoutes(r, projectHandler, taskHandler, auth)
	routes.TaskRoutes(r, taskHandler, auth)

	log.Printf("🚀 Server running on port %s", cfg.PORT)
	r.Run(":" + cfg.PORT)
}
