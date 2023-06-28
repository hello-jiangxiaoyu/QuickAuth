package service

import (
	"QuickAuth/pkg/model"
)

func (s *Service) GetUser(poolId, userName string) (*model.User, error) {
	var user model.User
	if err := s.db.Where("user_pool_id = ? AND username = ?", poolId, userName).
		First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *Service) ListUser(poolId string) ([]model.User, error) {
	var user []model.User
	if err := s.db.Where("user_Pool_id = ?", poolId).Find(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Service) CreateUser(u model.User) (*model.User, error) {
	if err := s.db.Create(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func (s *Service) ModifyUser(u model.User) (*model.User, error) {
	if err := s.db.Where("id = ? AND user_pool_id = ?", u.ID, u.UserPoolID).Save(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func (s *Service) DeleteUser(u model.User) (*model.User, error) {
	if err := s.db.Where("id = ? AND user_pool_id = ?", u.ID, u.UserPoolID).Delete(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}
