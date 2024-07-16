package controllers

import (
	"kaotonamae_back/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetGroups(c *gin.Context) {
	groups, err := services.GetGroups()
	if err != nil {
		logError(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"groups": groups,
	})
}

func GetGroupById(c *gin.Context) {
	groupId := c.Param("group_id")

	group, err := services.GetGroupById(groupId)
	if err != nil {
		logError(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"group": group,
	})
}

func GetGroupByUserId(c *gin.Context) {
	userId := c.Param("user_id")

	groups, err := services.GetGroupByUserId(userId)
	if err != nil {
		logError(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"groups": groups,
	})
}

func RegisterGroup(c *gin.Context) {
	var group services.GroupRequest

	if err := c.ShouldBindJSON(&group); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	res, err := services.RegisterGroup(group)
	if err != nil {
		logError(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"group": res,
	})
}