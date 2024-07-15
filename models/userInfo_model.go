package models

import (
	"time"
)

type UserInfo struct {
	UserId string `json:"user_id" gorm:"primaryKey"`
	UserLastName string `json:"user_last_name" gorm:"not null"`
	UserFirstName string `json:"user_first_name" gorm:"not null"`
	UserLastNameKana string `json:"user_last_name_kana" gorm:"not null"`
	UserFirstNameKana string `json:"user_first_name_kana" gorm:"not null"`
	Gender string `json:"gender" gorm:"not null"`
	Icon string `json:"icon" gorm:"not null"`
	BirthDate string `json:"birth_date" gorm:"not null"`
	Hobby string `json:"hobby" gorm:"not null"`
	Organization string `json:"organization"`
	HolidayActivity string `json:"holiday_activity"`
	Weakness string `json:"weakness"`
	FavoriteColor string `json:"favorite_color"`
	FavoriteAnimal string `json:"favorite_animal"`
	FavoritePlace string `json:"favorite_place"`
	Language string `json:"language"`
	Nickname string `json:"nickname"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

func GetAllUserInfos() (userInfos []UserInfo, err error) {
	err = DB.Debug().Find(&userInfos).Error

	return
}

func GetUserInfoById(userId string) (userInfo UserInfo, err error) {
	err = DB.Debug().Where("user_id = ?", userId).First(&userInfo).Error

	return
}

func RegisterUserInfo(userInfo UserInfo) (err error) {
	err = DB.Debug().Create(&userInfo).Error

	if err != nil {
		return err
	}

	return
}