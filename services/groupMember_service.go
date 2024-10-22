package services

import (
	"fmt"
	"kaotonamae_back/models"
)

type GroupMemberRequest struct {
	GroupId string `json:"group_id"`
	UserId string `json:"user_id"`
}

func GetGroupMembersByGroupId(groupId string) ([]string, error) {
	datas, err := models.GetGroupMembersByGroupId(groupId)
	if err != nil {
		return nil, fmt.Errorf("予期せぬエラーが発生しました: %v", err)
	}
	var userIds []string
	for _, data := range datas {
		userIds = append(userIds, data.UserId)
	}
	return userIds, nil
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

func RegisterGroupMembers(requests []GroupMemberRequest) ([]models.GroupMember, error) {
	var groupMembers []models.GroupMember
	for _, request := range requests {
		groupMember := models.GroupMember{
			GroupId: request.GroupId,
			UserId: request.UserId,
		}
		groupMembers = append(groupMembers, groupMember)
	}

	err := models.RegisterGroupMembers(groupMembers)
	if err != nil {
		return nil, fmt.Errorf("予期せぬエラーが発生しました: %v", err)
	}

	return groupMembers, nil
}

func DeleteGroupMember(groupId string, userId string) (models.GroupMember, error) {
	groupMember, err := models.DeleteGroupMember(groupId, userId)
	if err != nil {
		return models.GroupMember{},fmt.Errorf("予期せぬエラーが発生しました: %v", err)
	}

	return groupMember, nil
}


func DeleteGroupMembers(requests []GroupMemberRequest) ([]models.GroupMember, error) {
	var groupMembers []models.GroupMember
	for _, request := range requests {
		groupMember := models.GroupMember{
			GroupId: request.GroupId,
			UserId: request.UserId,
		}
		groupMembers = append(groupMembers, groupMember)
	}

	err := models.DeleteGroupMembers(groupMembers)
	if err != nil {
		return nil, fmt.Errorf("予期せぬエラーが発生しました: %v", err)
	}

	return groupMembers, nil
}