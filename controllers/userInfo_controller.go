package controllers

import (
	"kaotonamae_back/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserInfos(c *gin.Context) {
	userInfos, err := services.GetUserInfos()
	if err != nil {
		logError(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"userInfos": userInfos,
	})
}

func GetUserInfoById(c *gin.Context) {
	userId := c.Param("user_id")

	userInfo, err := services.GetUserInfoById(userId)
	if err != nil {
		logError(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"userInfo": userInfo,
	})
}

func RegisterUserInfo(c *gin.Context) {
	var userInfo services.UserInfoRequest

	if err := c.ShouldBindJSON(&userInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	res, err := services.RegisterUserInfo(userInfo)
	if err != nil {
		logError(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"userInfo": res})
}

func UpdateUserInfoById(c *gin.Context) {
	userId := c.Param("user_id")

	var userInfo services.UserInfoRequest

	if err := c.ShouldBindJSON(&userInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	res, err := services.UpdateUserInfoById(userId, userInfo)
	if err != nil {
		logError(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"userInfo": res,
	})
}

