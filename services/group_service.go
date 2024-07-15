package services

import (
	"fmt"
	"kaotonamae_back/models"
)

type GroupRequest struct {
	GroupId string `json:"group_id"`
	UserId string `json:"user_id"`
	GroupName string `json:"group_name"`
	GroupDescription string `json:"group_description"`
	GroupIcon string `json:"group_icon"`
}

func GetGroups() ([]models.Group, error) {
	datas, err := models.GetAllGroups()
	if err != nil {
		return nil, fmt.Errorf("予期せぬエラーが発生しました: %v", err)
	}
	return datas, nil
}

func GetGroupById(groupId string) (models.Group, error) {
	data, err := models.GetGroupById(groupId)
	if err != nil {
		return models.Group{}, fmt.Errorf("予期せぬエラーが発生しました: %v", err)
	}
	return data, nil
}

func RegisterGroup(request GroupRequest) (models.Group, error) {
	group := models.Group{
		GroupId: request.GroupId,
		UserId: request.UserId,
		GroupName: request.GroupName,
		GroupDescription: request.GroupDescription,
		GroupIcon: request.GroupIcon,
	}

	err := models.RegisterGroup(group)
	if err != nil {
		return models.Group{}, fmt.Errorf("予期せぬエラーが発生しました: %v", err)
	}

	return group, nil
}