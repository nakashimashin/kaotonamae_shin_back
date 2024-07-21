package controllers

import (
	"kaotonamae_back/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetFriends(c *gin.Context) {
	friends, err := services.GetFriends()
	if err != nil {
		logError(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"friends": friends,
	})
}

func GetFriendsByUserId(c *gin.Context) {
	userId := c.Param("user_id")

	friends, err := services.GetFriendsByUserId(userId)
	if err != nil {
		logError(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"friends": friends,
	})
}

func RegisterFriend(c *gin.Context) {
	var friend services.FriendRequest
	if err := c.ShouldBindJSON(&friend); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	res, err := services.RegisterFriend(friend)
	if err != nil {
		logError(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"friend": res,
	})
}