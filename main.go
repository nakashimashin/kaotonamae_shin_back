package main

import (
	"net/http"
	"os"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Result struct {
	ID int
	Name string
}

func main() {
	engine := gin.Default()
	engine.GET("/", func(c *gin.Context) {
		dsn := fmt.Sprintf("%s:%s@tcp(kaotonamae_db:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local"
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_DATABASE"),
		)
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}

		var result Result
		db.First(&result)

		c.JSON(http.StatusOK, gin.H{
			"test" : result.Name,
		})
	})
	engine.Run(":8080")
}
