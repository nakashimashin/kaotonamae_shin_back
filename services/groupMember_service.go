package services

import (
	"fmt"
	"kaotonamae_back/models"
)

type GroupMemberRequest struct {
	GroupId string `json:"group_id"`
	UserId string `json:"user_id"`
}

func RegisterGroupMember(request GroupMemberRequest) (models.GroupMember, error) {
	groupMember := models.GroupMember{
		GroupId: request.GroupId,
		UserId: request.UserId,
	}

	err := models.RegisterGroupMember(groupMember)
	if err != nil {
		return models.GroupMember{},fmt.Errorf("予期せぬエラーが発生しました: %v", err)
	}

	return groupMember, nil
}