package models

import (
	"time"
)

type Auth struct {
	UserId string `json:"user_id" gorm:"primaryKey"`
	Email string `json:"email" gorm:"not null; unique"`
	Password string `json:"password" gorm:"not null"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

func GetAllAuths() (auths []Auth, err error) {
	err = DB.Debug().Find(&auths).Error

	return
}