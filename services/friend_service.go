package services

import (
	"fmt"
	"kaotonamae_back/models"

	// "github.com/google/uuid"
)

type FriendRequest struct {
	FriendId string `json:"friend_id"`
	UserId string `json:"user_id"`
}

func GetFriends() ([]models.Friend, error) {
	datas, err := models.GetAllFriends()
	if err != nil {
		return nil, fmt.Errorf("予期せぬエラーが発生しました: %v", err)
	}
	return datas, nil
}

func GetFriendById(friendId string) (models.Friend, error) {
	data, err := models.GetFriendById(friendId)
	if err != nil {
		return models.Friend{}, fmt.Errorf("予期せぬエラーが発生しました: %v", err)
	}
	return data, nil
}

func GetFriendByUserId(userId string) ([]models.Friend, error) {
	datas, err := models.GetFriendByUserId(userId)
	if err != nil {
		return nil, fmt.Errorf("予期せぬエラーが発生しました: %v", err)
	}
	return datas, nil
}