package controllers

import (
	"kaotonamae_back/services"
	"net/http"
	"log"

	"github.com/gin-gonic/gin"
)

func GetAuths(c *gin.Context) {

	auths, err := services.GetAuths()
	if err != nil {
		logError(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"auths": auths,
	})
}

func RegisterAuth(c *gin.Context) {
	var auth services.AuthRequest

	if err := c.ShouldBindJSON(&auth); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	res, err := services.RegisterAuth(auth)
	if err != nil {
		logError(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"auth": res})
}

func logError(err error) {
	log.Printf("error occurred: %v",err)
}

