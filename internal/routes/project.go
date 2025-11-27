package routes

import (
	"gin/internal/domain/project"
	"gin/internal/domain/task"

	"github.com/gin-gonic/gin"
)

func ProjectRoutes(r *gin.Engine, hp *project.ProjectHandler, ht *task.TaskHandler, auth gin.HandlerFunc) {
	p := r.Group("/projects")
	p.Use(auth)

	p.GET("/:id", hp.FindByID)
	p.PATCH("/:id", hp.Update)
	p.DELETE("/:id", hp.Delete)

	p.POST("/:id/tasks", ht.Create)       // create task
	p.GET("/:id/tasks", ht.FindByProject) // find task by project
}
