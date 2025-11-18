package routes

import (
	"gin/internal/domain/user"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine, h *user.UserHandler) {
	r.POST("/register", h.Register)
	r.POST("/login", h.Login)
	r.POST("/refresh", h.Refresh)
}
