package routes

import (
	"gin/internal/domain/task"

	"github.com/gin-gonic/gin"
)

func TaskRoutes(r *gin.Engine, h *task.TaskHandler, auth gin.HandlerFunc) {
	t := r.Group("/tasks")
	t.Use(auth)

	t.GET("/:id", h.FindByID)
	t.PATCH("/:id", h.Update)
	t.DELETE("/:id", h.Delete)
}
