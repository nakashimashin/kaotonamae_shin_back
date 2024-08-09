package models

import (
	"time"
)

type GroupMember struct {
	GroupId string `json:"group_id" gorm:"primaryKey"`
	UserId string `json:"user_id" gorm:"primaryKey"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

func RegisterGroupMember(groupMember GroupMember) (err error) {
	err = DB.Debug().Create(&groupMember).Error

	if err != nil {
		return err
	}

	return
}