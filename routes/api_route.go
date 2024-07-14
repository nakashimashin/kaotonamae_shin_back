package routes

import (
	"kaotonamae_back/controllers"

	"github.com/gin-gonic/gin"
)

func GetApiRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.GET("/", controllers.GetAuths)
		}
	}

	return r
}