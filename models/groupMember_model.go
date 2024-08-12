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

func GetGroupMembersByGroupId(groupId string) (groupMembers []GroupMember, err error) {
	err = DB.Debug().Where("group_id = ?", groupId).Find(&groupMembers).Error

	return
}

func RegisterGroupMember(groupMember GroupMember) (err error) {
	err = DB.Debug().Create(&groupMember).Error

	if err != nil {
		return err
	}

	return
}

func RegisterGroupMembers(groupMembers []GroupMember) (err error) {
	err = DB.Debug().Create(&groupMembers).Error

	if err != nil {
		return err
	}

	return
}

func DeleteGroupMember(groupId string, userId string) (groupMember GroupMember, err error) {
	err = DB.Debug().Where("group_id = ? AND user_id = ?", groupId, userId).Delete(&groupMember).Error

	return
}