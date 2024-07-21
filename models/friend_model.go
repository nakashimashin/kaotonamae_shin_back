package models

import (
	"time"
)

type Friend struct {
	UserId1 string `json:"user_id1" gorm:"primaryKey"`
	UserId2 string `json:"user_id2" gorm:"primaryKey"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

func GetAllFriends() (friends []Friend, err error) {
	err = DB.Debug().Find(&friends).Error

	return
}

func GetFriendsByUserId(userId string) (friends []Friend, err error) {
	err = DB.Debug().Where("user_id1 = ? OR user_id2 = ?", userId, userId).Find(&friends).Error

	return
}

func RegisterFriend(friend Friend) (err error) {
	err = DB.Debug().Create(&friend).Error

	if err != nil {
		return err
	}

	return
}