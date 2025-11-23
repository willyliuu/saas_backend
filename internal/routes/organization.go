package routes

import (
	"gin/internal/domain/organization"

	"github.com/gin-gonic/gin"
)

func OrganizationRoutes(r *gin.Engine, h *organization.OrganizationHandler, auth gin.HandlerFunc) {
	o := r.Group("/organizations")
	o.Use(auth)

	o.POST("/", h.Create)
	o.GET("/", h.FindAll)
	o.GET("/:id", h.FindByID)
}
