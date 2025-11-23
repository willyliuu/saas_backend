package routes

import (
	"gin/internal/domain/user"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine, h *user.UserHandler, auth gin.HandlerFunc) {
	u := r.Group("/users")
	u.Use(auth)

	u.GET("/me", h.Me)
}
