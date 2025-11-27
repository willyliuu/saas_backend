package routes

import (
	"gin/internal/domain/invite"
	"gin/internal/domain/organization"
	"gin/internal/domain/project"

	"github.com/gin-gonic/gin"
)

func OrganizationRoutes(r *gin.Engine, ho *organization.OrganizationHandler, hi *invite.InviteHandler, hp *project.ProjectHandler, auth gin.HandlerFunc) {
	o := r.Group("/organizations")
	o.Use(auth)

	o.POST("/", ho.Create)
	o.GET("/", ho.FindAll)
	o.GET("/:id", ho.FindByID)
	o.PATCH("/:id", ho.Update)
	o.DELETE("/:id", ho.Delete)

	o.POST("/:id/invites", hi.InviteMember) // send invite
	o.POST("/:id/projects", hp.Create)      // create project
	o.GET("/:id/projects", hp.FindByOrgID)  // find all project by orgID
}
