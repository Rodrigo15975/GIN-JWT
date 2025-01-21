package config

import (
	"db7/routers"

	"github.com/gin-gonic/gin"
)

func InitApp() {
    
    r := gin.Default()
    routers.UserRoutes(r)
	r.Run(":4000")
}
