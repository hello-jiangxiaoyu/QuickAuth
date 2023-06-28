package service

import (
	"QuickAuth/internal/global"
	"QuickAuth/pkg/models"
)

func (s *Service) GetClientById(id string) (*models.Client, error) {
	var client models.Client
	if err := global.DB.Where("id = ?", id).First(&client).Error; err != nil {
		return nil, err
	}
	return &client, nil
}
