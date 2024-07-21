package models

import (
	"time"
)

type Friend struct {
	FriendId string `json:"friend_id" gorm:"primaryKey"`
	UserId string `json:"user_id" gorm:"not null"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

func GetAllFriends() (friends []Friend, err error) {
	err = DB.Debug().Find(&friends).Error

	return
}

func GetFriendById(friendId string) (friend Friend, err error) {
	err = DB.Debug().Where("friend_id = ?", friendId).First(&friend).Error

	return
}

func GetFriendByUserId(userId string) (friends []Friend, err error) {
	err = DB.Debug().Where("user_id = ?", userId).Find(&friends).Error

	return
}