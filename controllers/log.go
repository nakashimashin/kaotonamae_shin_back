package controllers

import (
	"log"
)

func logError(err error) {
	log.Printf("error occurred: %v",err)
}