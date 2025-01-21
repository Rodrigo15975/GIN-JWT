package controllers

import (
	"db7/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	users, err := services.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch users in services",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}
