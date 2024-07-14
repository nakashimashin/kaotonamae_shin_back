package main

import (
	"kaotonamae_back/models"
	"kaotonamae_back/routes"
)


func main() {
	models.Migrate()

	r := routes.GetApiRouter()
	r.Run(":8080")
}