package services

import (
	"fmt"
	"kaotonamae_back/models"

)

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