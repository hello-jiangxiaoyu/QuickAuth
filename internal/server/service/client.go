package service

import (
	"QuickAuth/internal/global"
	"QuickAuth/internal/model"
)

func GetClientById(id string) (*model.Client, error) {
	var client model.Client
	if err := global.DB.Where("id = ?", id).First(&client).Error; err != nil {
		return nil, err
	}
	return &client, nil
}
