package invite

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type InviteHandler struct {
	Service *InviteService
}

func (h *InviteHandler) InviteMember(c *gin.Context) {
	orgIDParam := c.Param("id")
	orgID, err := strconv.Atoi(orgIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var dto CreateInviteDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	invite, err := h.Service.InviteMember(uint(orgID), dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "invite sent",
		"token":   invite.Token,
	})
}

func (h *InviteHandler) AcceptInvite(c *gin.Context) {
	token := c.Param("token")
	userID := c.GetUint("user_id")

	// check token exists
	invite, err := h.Service.GetInviteByToken(token)
	if err != nil || invite.Status != "pending" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid token or expired invite"})
		return
	}

	userEmail := c.GetString("user_email")
	if invite.Email != userEmail {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invite is not for this user"})
		return
	}

	userOrg, err := h.Service.AcceptInvite(userID, invite)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      "invited accepted",
		"organization": userOrg,
	})
}
