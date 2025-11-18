package user

import "github.com/gin-gonic/gin"

type UserHandler struct {
	Service *UserService
}

func (h *UserHandler) Register(c *gin.Context) {
	var dto RegisterDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := h.Service.Register(dto); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{"message": "registered"})
}

func (h *UserHandler) Login(c *gin.Context) {
	var dto LoginDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	accessToken, refreshToken, err := h.Service.Login(dto)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}

	c.SetCookie("refresh_token", refreshToken, 3600*24*7, "/", "", true, true)

	c.JSON(200, gin.H{
		"access_token": accessToken,
	})
}

func (h *UserHandler) Refresh(c *gin.Context) {
	cookie, err := c.Cookie("refresh_token")
	if err != nil {
		c.JSON(400, gin.H{"error": "no refresh token"})
	}

	accessToken, refreshToken, err := h.Service.Refresh(cookie)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}

	c.SetCookie("refresh_token", refreshToken, 3600*24*7, "/", "", true, true)

	c.JSON(200, gin.H{
		"access_token": accessToken,
	})
}
