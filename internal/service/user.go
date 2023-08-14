package service

import (
	"QuickAuth/internal/model"
	"QuickAuth/pkg/utils"
	"github.com/pkg/errors"
)

func (s *Service) GetUserByName(poolId int64, userName string) (*model.User, error) {
	var user model.User
	if err := s.db.Where("user_pool_id = ? AND username = ?", poolId, userName).
		First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *Service) GetUserById(poolId int64, userId string) (*model.User, error) {
	var user model.User
	if err := s.db.Select("id", "username", "display_name", "email", "phone").
		Where("id = ? AND user_pool_id = ?", userId, poolId).
		First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *Service) ListUser(poolId int64) ([]model.User, error) {
	var user []model.User
	if err := s.db.Select("id", "username", "display_name", "email", "phone").
		Where("user_Pool_id = ?", poolId).Find(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) CreateUser(u *model.User) (*model.User, error) {
	if _, err := s.GetUserPool(u.UserPoolID); err != nil {
		return nil, errors.Wrap(err, "no such user pool")
	}

	u.ID = utils.GetNoLineUUID()
	if err := s.db.Create(u).Error; err != nil {
		return nil, err
	}
	u.Password = ""
	return u, nil
}

func (s *Service) ModifyUser(userId string, u *model.User) error {
	if _, err := s.GetUserPool(u.UserPoolID); err != nil {
		return errors.Wrap(err, "no such user pool")
	}

	if err := s.db.Select("display_name", "email", "phone").
		Where("id = ? AND user_pool_id = ?", userId, u.UserPoolID).
		Updates(u).Error; err != nil {
		return err
	}
	return nil
}

func (s *Service) DeleteUser(poolId int64, userId string) error {
	if err := s.db.Where("id = ? AND user_pool_id = ?", userId, poolId).Delete(&model.User{}).Error; err != nil {
		return err
	}
	return nil
}

// ====================== user pool ======================

func (s *Service) GetUserPool(poolId int64) (*model.UserPool, error) {
	var pool model.UserPool
	if err := s.db.Where("id = ?", poolId).First(&pool).Error; err != nil {
		return nil, err
	}
	return &pool, nil
}

func (s *Service) ListUserPool() ([]model.UserPool, error) {
	var pool []model.UserPool
	if err := s.db.Select("id", "name", "describe", "created_at").Find(&pool).Error; err != nil {
		return nil, err
	}
	return pool, nil
}

func (s *Service) CreateUserPool(pool *model.UserPool) (*model.UserPool, error) {
	if err := s.db.Create(pool).Error; err != nil {
		return nil, err
	}
	return pool, nil
}

func (s *Service) ModifyUserPool(poolId int64, pool *model.UserPool) error {
	if err := s.db.Where("id = ?", poolId).Updates(pool).Error; err != nil {
		return err
	}
	return nil
}

func (s *Service) DeleteUserPool(poolId int64) error {
	if err := s.db.Where("id = ?", poolId).Delete(&model.UserPool{}).Error; err != nil {
		return err
	}
	return nil
}
