package services

import (
	"fmt"
	"kaotonamae_back/models"
)

type UserInfoRequest struct {
	UserId string `json:"user_id"`
	UserLastName string `json:"user_last_name"`
	UserFirstName string `json:"user_first_name"`
	UserLastNameKana string `json:"user_last_name_kana"`
	UserFirstNameKana string `json:"user_first_name_kana"`
	Gender string `json:"gender"`
	Icon string `json:"icon"`
	BirthDate string `json:"birth_date"`
	Hobby string `json:"hobby"`
	Organization string `json:"organization"`
	HolidayActivity string `json:"holiday_activity"`
	Weakness string `json:"weakness"`
	FavoriteColor string `json:"favorite_color"`
	FavoriteAnimal string `json:"favorite_animal"`
	FavoritePlace string `json:"favorite_place"`
	Language string `json:"language"`
	Nickname string `json:"nickname"`
}

func GetUserInfos() ([]models.UserInfo, error) {
	datas, err := models.GetAllUserInfos()
	if err != nil {
		return nil, fmt.Errorf("予期せぬエラーが発生しました: %v", err)
	}
	return datas, nil
}

func GetUserInfoById(userId string) (models.UserInfo, error) {
	data, err := models.GetUserInfoById(userId)
	if err != nil {
		return models.UserInfo{}, fmt.Errorf("予期せぬエラーが発生しました: %v", err)
	}
	return data, nil
}

func RegisterUserInfo(request UserInfoRequest) (models.UserInfo, error) {
	userInfo := models.UserInfo{
		UserId: request.UserId,
		UserLastName: request.UserLastName,
		UserFirstName: request.UserFirstName,
		UserLastNameKana: request.UserLastNameKana,
		UserFirstNameKana: request.UserFirstNameKana,
		Gender: request.Gender,
		Icon: request.Icon,
		BirthDate: request.BirthDate,
		Hobby: request.Hobby,
		Organization: request.Organization,
		HolidayActivity: request.HolidayActivity,
		Weakness: request.Weakness,
		FavoriteColor: request.FavoriteColor,
		FavoriteAnimal: request.FavoriteAnimal,
		FavoritePlace: request.FavoritePlace,
		Language: request.Language,
		Nickname: request.Nickname,
	}
	err := models.RegisterUserInfo(userInfo)
	if err != nil {
		return models.UserInfo{}, fmt.Errorf("予期せぬエラーが発生しました: %v", err)
	}
	return userInfo, nil
}