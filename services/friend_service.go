package services

import (
	"fmt"
	"kaotonamae_back/models"

)

type FriendRequest struct {
	UserId1 string `json:"user_id1"`
	UserId2 string `json:"user_id2"`
}

func GetFriends() ([]models.Friend, error) {
	datas, err := models.GetAllFriends()
	if err != nil {
		return nil, fmt.Errorf("予期せぬエラーが発生しました: %v", err)
	}
	return datas, nil
}

func GetFriendsByUserId(userId string) ([]models.Friend, error) {
	datas, err := models.GetFriendsByUserId(userId)
	if err != nil {
		return nil, fmt.Errorf("予期せぬエラーが発生しました: %v", err)
	}
	var friendIds []string
	for _, data := range datas {
		if data.UserId1 == userId {
			friendIds = append(friendIds, data.UserId2)
		} else {
			friendIds = append(friendIds, data.UserId1)
		}
	}
	return friendIds, nil
}
