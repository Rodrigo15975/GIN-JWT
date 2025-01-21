package services

import (
	"db7/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllUsers() ([]models.User, error) {
	users := []models.User{}

	if len(users) == 0 {
		return nil, gin.Error{
			Err:  http.ErrNoLocation,
			Type: gin.ErrorTypePublic,
		}
	}

	return users, nil
}
