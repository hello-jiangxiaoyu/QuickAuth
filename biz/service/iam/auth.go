package iam

import (
	"QuickAuth/biz/endpoint/model"
	"QuickAuth/pkg/global"
	"errors"
	"gorm.io/gorm"
)

func IsOperationAllow(tenantId int64, resId int64, nodeId int64, operationId int64, userId any) (bool, error) {
	var roles []int64
	if err := global.Db().Model(model.ResourceRole{}).Select("id").
		Where("tenant_id = ? AND resource_id = ? AND operation_id = ?", tenantId, resId, operationId).Find(&roles).Error; err != nil {
		return false, err
	}

	err := global.Db().Where("user_id = ? AND tenant_id = ? AND node_id = ? AND resource_id = ? AND role_id in ?",
		userId, tenantId, nodeId, resId, roles).First(model.ResourceUserRole{}).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return true, nil
}

func IsJSONOperationAllow(tenantId int64, resId int64, path string, operationId int64, userId any) (bool, error) {
	var roles []int64
	if err := global.Db().Model(model.ResourceRole{}).Select("id").
		Where("tenant_id = ? AND resource_id = ? AND operation_id = ?", tenantId, resId, operationId).Find(&roles).Error; err != nil {
		return false, err
	}

	err := global.Db().Where("user_id = ? AND tenant_id = ? AND json_path = ? AND resource_id = ? AND role_id in ?",
		userId, tenantId, path, resId, roles).First(model.ResourceJSONUserRole{}).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	} else if err != nil {
		return false, err
	}

	return true, nil
}

// ListResourceOperationNodes 获取拥有某个操作权限的node列表
func ListResourceOperationNodes(tenantId int64, resId int64, paraId int64, operationId int64, userId string) ([]model.ResourceUserRole, error) {
	var data []model.ResourceUserRole
	if err := global.Db().Table(model.TableNameResourceUserRole+" as rur").
		Select("rur.id", "rur.tenant_id", "rur.user_id", "rur.role_id", "rur.resource_id", "rur.node_id").
		Joins(model.TableNameResourceNode+" as n ON n.id = rur.node_id").
		Joins(model.TableNameResourceRoleOperation+" as ro ON ro.id = rur.role_id").
		Where("rur.tenant_id = ? AND rur.resource_id = ? AND rur.user_id = ? AND ro.operation_id = ? AND n.parent = ?",
			tenantId, resId, userId, operationId, paraId).Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

// ListJSONResourceOperationNodes 获取拥有某个操作权限的整个json结构
func ListJSONResourceOperationNodes(tenantId int64, resId int64, operationId int64, userId string) ([]model.ResourceJSONUserRole, error) {
	var data []model.ResourceJSONUserRole
	if err := global.Db().Table(model.TableNameResourceJSONUserRole+" as rj").
		Select("rj.id", "rj.tenant_id", "rj.resource_id", "rj.json_path", "rj.user_id", "rj.role_id").
		Joins(model.TableNameResourceRoleOperation+" as ro ON ro.id = rj.role_id").
		Where("rj.tenant_id = ? AND rj.resource_id = ? AND rj.user_id = ? AND ro.operation_id = ?", tenantId, resId, userId, operationId).
		Find(&data).Error; err != nil {
		return nil, err
	}

	// todo: 根据path拼接原始json结构

	return data, nil
}
