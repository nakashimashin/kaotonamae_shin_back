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

func logError(err error) {
	log.Printf("error occurred: %v",err)
}

