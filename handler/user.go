package handler

import (
	"Mrkonxyz/github.com/model"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	err := h.Service.Repo.CreateUser(&user)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create user: " + err.Error()})
		return
	}

	c.JSON(201, gin.H{"message": "User created successfully"})
}

func (h *Handler) GetUserByUsername(c *gin.Context) {

	username := c.Param("username")

	result, err := h.Service.Repo.GetUserByUsername(username)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to get user"})
		return
	}

	c.JSON(200, gin.H{"user": result})
}
