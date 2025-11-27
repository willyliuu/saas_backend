package routes

import (
	"gin/internal/domain/invite"

	"github.com/gin-gonic/gin"
)

func InviteRoutes(r *gin.Engine, h *invite.InviteHandler, auth gin.HandlerFunc) {
	i := r.Group("/invites")

	i.Use(auth)

	i.POST("/:token/accept", h.AcceptInvite) // accept invite
}
