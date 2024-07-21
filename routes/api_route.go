package routes

import (
	"kaotonamae_back/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func GetApiRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Origin", "Content-Length", "Content-Type"},
		ExposeHeaders: []string{"Content-Length"},
	}))

	v1 := r.Group("/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.GET("/", controllers.GetAuths)
			auth.POST("/", controllers.RegisterAuth)
		}

		userInfo := v1.Group("/userInfo")
		{
			userInfo.GET("/", controllers.GetUserInfos)
			userInfo.GET("/:user_id", controllers.GetUserInfoById)
			userInfo.POST("/", controllers.RegisterUserInfo)
		}

		group := v1.Group("/group")
		{
			group.GET("/", controllers.GetGroups)
			group.GET("/:group_id", controllers.GetGroupById)
			group.GET("/user/:user_id", controllers.GetGroupByUserId)
			group.POST("/", controllers.RegisterGroup)
			group.PUT("/:group_id", controllers.UpdateGroupById)
		}

		friend := v1.Group("/friend")
		{
			friend.GET("/", controllers.GetFriends)
			friend.GET("/:friend_id", controllers.GetFriendById)
			friend.GET("/user/:user_id", controllers.GetFriendByUserId)
		}
	}

	return r
}