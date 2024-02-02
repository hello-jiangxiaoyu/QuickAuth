package admin

import (
	"QuickAuth/biz/endpoint/model"
	"QuickAuth/pkg/global"
	"QuickAuth/pkg/utils"
)

func GetUserByName(poolId int64, userName string) (*model.User, error) {
	var user model.User
	if err := global.Db().Where("user_pool_id = ? AND username = ?", poolId, userName).
		First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserById(poolId int64, userId string) (*model.User, error) {
	var user model.User
	if err := global.Db().Select("id", "username", "display_name", "email", "phone").
		Where("id = ? AND user_pool_id = ?", userId, poolId).
		First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func ListUser(poolId int64) ([]model.User, error) {
	var user []model.User
	if err := global.Db().Select("id", "username", "display_name", "email", "phone").
		Where("user_Pool_id = ?", poolId).Find(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func CreateUser(u *model.User) (*model.User, error) {
	if _, err := GetUserPool(u.UserPoolID); err != nil {
		return nil, utils.WithMessage(err, "no such user pool")
	}

	u.ID = utils.GetNoLineUUID()
	if err := global.Db().Create(u).Error; err != nil {
		return nil, err
	}
	u.Password = ""
	return u, nil
}

func ModifyUser(userId string, u *model.User) error {
	if _, err := GetUserPool(u.UserPoolID); err != nil {
		return utils.WithMessage(err, "no such user pool")
	}

	if err := global.Db().Select("display_name", "email", "phone").
		Where("id = ? AND user_pool_id = ?", userId, u.UserPoolID).
		Updates(u).Error; err != nil {
		return err
	}
	return nil
}

func DeleteUser(poolId int64, userId string) error {
	if err := global.Db().Where("id = ? AND user_pool_id = ?", userId, poolId).Delete(&model.User{}).Error; err != nil {
		return err
	}
	return nil
}

// ====================== user pool ======================

func GetUserPool(poolId int64) (*model.UserPool, error) {
	var pool model.UserPool
	if err := global.Db().Where("id = ?", poolId).First(&pool).Error; err != nil {
		return nil, err
	}
	return &pool, nil
}

func ListUserPool() ([]model.UserPool, error) {
	var pool []model.UserPool
	if err := global.Db().Select("id", "name", "describe", "created_at").Find(&pool).Error; err != nil {
		return nil, err
	}
	return pool, nil
}

func CreateUserPool(pool *model.UserPool) (*model.UserPool, error) {
	if err := global.Db().Create(pool).Error; err != nil {
		return nil, err
	}
	return pool, nil
}

func ModifyUserPool(poolId int64, pool *model.UserPool) error {
	if err := global.Db().Where("id = ?", poolId).Updates(pool).Error; err != nil {
		return err
	}
	return nil
}

func DeleteUserPool(poolId int64) error {
	if err := global.Db().Where("id = ?", poolId).Delete(&model.UserPool{}).Error; err != nil {
		return err
	}
	return nil
}
