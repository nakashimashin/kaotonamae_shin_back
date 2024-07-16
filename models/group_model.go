package models

import (
	"time"
)

type Group struct {
	GroupId string `json:"group_id" gorm:"primaryKey"`
	UserId string `json:"user_id" gorm:"not null"`
	GroupName string `json:"group_name" gorm:"not null"`
	GroupDescription string `json:"group_description" gorm:"not null"`
	GroupIcon string `json:"group_icon" gorm:"not null"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

func GetAllGroups() (groups []Group, err error) {
	err = DB.Debug().Find(&groups).Error

	return
}

func GetGroupById(groupId string) (group Group, err error) {
	err = DB.Debug().Where("group_id = ?", groupId).First(&group).Error

	return
}

func GetGroupByUserId(userId string) (groups []Group, err error) {
	err = DB.Debug().Where("user_id = ?", userId).Find(&groups).Error

	return
}

func RegisterGroup(group Group) (err error) {
	err = DB.Debug().Create(&group).Error

	if err != nil {
		return err
	}

	return
}