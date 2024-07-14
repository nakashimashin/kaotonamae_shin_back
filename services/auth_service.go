package services

import (
	"fmt"
	"kaotonamae_back/models"

)

type AuthRequest struct {
	UserId string `json:"user_id"`
	Email string `json:"email"`
	Password string `json:"password"`
}

func GetAuths() ([]string, error) {
	datas, err := models.GetAllAuths()
	if err != nil {
		return nil, fmt.Errorf("予期せぬエラーが発生しました: %v", err)
	}
	var auths []string
	for _, data := range datas {
		auths = append(auths, data.Email)
	}
	return auths, nil
}

func RegisterAuth(request AuthRequest) (auth models.Auth, err error) {
	auth = models.Auth{
		UserId: request.UserId,
		Email: request.Email,
		Password: request.Password,
	}
	err = models.RegisterAuth(auth)
	if err != nil {
		return auth, fmt.Errorf("予期せぬエラーが発生しました: %v", err)
	}
	return auth, nil
}