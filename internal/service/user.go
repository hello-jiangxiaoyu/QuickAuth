package service

import (
	"QuickAuth/pkg/model"
	"QuickAuth/pkg/tools/utils"
)

func (s *Service) GetUserByName(poolId, userName string) (*model.User, error) {
	var user model.User
	if err := s.db.Where("user_pool_id = ? AND username = ?", poolId, userName).
		First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *Service) GetUserById(poolId, userId string) (*model.User, error) {
	var user model.User
	if err := s.db.Where("id = ? AND user_pool_id = ?", poolId, userId).
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

func (s *Service) ModifyUser(u model.User) error {
	if err := s.db.Select("display_name", "email", "phone").
		Where("id = ? AND user_pool_id = ?", u.ID, u.UserPoolID).
		Save(&u).Error; err != nil {
		return err
	}
	return nil
}

func (s *Service) DeleteUser(poolId, userId string) error {
	if err := s.db.Where("id = ? AND user_pool_id = ?", userId, poolId).Delete(&model.UserPool{}).Error; err != nil {
		return err
	}
	return nil
}

// ====================== user pool ======================

func (s *Service) GetUserPool(poolId string) (*model.UserPool, error) {
	var pool model.UserPool
	if err := s.db.Where("id = ?", poolId).
		First(&pool).Error; err != nil {
		return nil, err
	}
	return &pool, nil
}

func (s *Service) ListUserPool() ([]model.UserPool, error) {
	var pool []model.UserPool
	if err := s.db.Select("id", "name").Find(&pool).Error; err != nil {
		return nil, err
	}
	return pool, nil
}

func (s *Service) CreateUserPool(pool model.UserPool) (*model.UserPool, error) {
	pool.ID = utils.GetNoLineUUID()
	if err := s.db.Create(&pool).Error; err != nil {
		return nil, err
	}
	return &pool, nil
}

func (s *Service) ModifyUserPool(pool model.UserPool) error {
	if err := s.db.Where("id = ?", pool.ID).Save(&pool).Error; err != nil {
		return err
	}
	return nil
}

func (s *Service) DeleteUserPool(poolId string) error {
	if err := s.db.Where("id = ? AND user_pool_id = ?", poolId).Delete(&model.UserPool{}).Error; err != nil {
		return err
	}
	return nil
}