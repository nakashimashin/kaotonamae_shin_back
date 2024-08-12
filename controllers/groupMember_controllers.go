package controllers

import (
	"kaotonamae_back/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetGroupMembersByGroupId(c *gin.Context) {
	groupId := c.Param("group_id")

	groupMembers, err := services.GetGroupMembersByGroupId(groupId)
	if err != nil {
		logError(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"groupMembers": groupMembers,
	})
}

func RegisterGroupMember(c *gin.Context) {
	var groupMember services.GroupMemberRequest
	if err := c.ShouldBindJSON(&groupMember); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	res, err := services.RegisterGroupMember(groupMember)
	if err != nil {
		logError(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"groupMember": res,
	})
}

func DeleteGroupMember(c *gin.Context) {
	groupId := c.Param("group_id")
	userId := c.Param("user_id")

	res, err := services.DeleteGroupMember(groupId, userId)
	if err != nil {
		logError(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"groupMember": res,
	})
}