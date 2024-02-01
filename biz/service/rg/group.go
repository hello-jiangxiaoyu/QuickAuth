package rg

import (
	"QuickAuth/biz/endpoint/model"
	"QuickAuth/pkg/global"
	"gorm.io/gorm"
)

func GetResourceGroupList(tenantId int64) ([]model.ResourceGroup, error) {
	var groupList []model.ResourceGroup
	if err := global.Db().Where("tenant_id = ?", tenantId).Find(&groupList).Error; err != nil {
		return nil, err
	}
	return groupList, nil
}

func GetResourceGroup(tenantId, groupId int64) (*model.ResourceGroup, error) {
	var group model.ResourceGroup
	if err := global.Db().Where("id = ? AND tenant_id = ?", groupId, tenantId).First(&group).Error; err != nil {
		return nil, err
	}
	return &group, nil
}

func CreateResourceGroup(tenantId, uid int64, name string, des string) (*model.ResourceGroup, error) {
	group := model.ResourceGroup{
		Id:          uid,
		TenantId:    tenantId,
		Name:        name,
		Description: des,
	}
	if err := global.Db().Create(&group).Error; err != nil {
		return nil, err
	}
	return &group, nil
}

func UpdateResourceGroup(tenantId, groupId int64, name, des string) error {
	if err := global.Db().Where("id = ? AND tenant_id = ?", groupId, tenantId).
		Model(&model.ResourceGroup{}).UpdateColumns(map[string]any{"name": name, "description": des}).Error; err != nil {
		return err
	}
	return nil
}

func DeleteResourceGroup(tenantId, groupId int64) error {
	if err := global.Db().Where("id = ? AND tenant_id = ?", groupId, tenantId).
		First(&model.ResourceGroup{}).Error; err != nil {
		return err
	}
	var roles []string
	if err := global.Db().Select("id").Where("group_id = ? AND tenant_id = ?", groupId, tenantId).
		Model(&model.ResourceGroupRole{}).Find(&roles).Error; err != nil {
		return err
	}
	deleteList := []any{
		model.ResourceGroupUser{},
		model.ResourceGroupAction{},
		model.ResourceGroupRole{},
		model.ResourceGroupResource{},
	}
	return global.Db().Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("role_id in ? AND tenant_id = ?", roles, tenantId).Delete(&model.ResourceGroupRoleAction{}).Error; err != nil {
			return err
		}
		for _, m := range deleteList {
			if err := tx.Where("group_id = ? AND tenant_id = ?", groupId, tenantId).
				Delete(&m).Error; err != nil {
				return err
			}
		}
		if err := tx.Where("id = ? AND tenant_id = ?", groupId, tenantId).
			Delete(&model.ResourceGroup{}).Error; err != nil {
			return err
		}
		return nil
	})
}

func GetResourceGroupResourceList(tenantId, groupId int64) ([]model.ResourceGroupResource, error) {
	var res []model.ResourceGroupResource
	if err := global.Db().Where("group_id = ? AND tenant_id = ?", groupId, tenantId).Find(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}

func GetResourceGroupResource(tenantId, groupId, resourceId int64) (*model.ResourceGroupResource, error) {
	var res model.ResourceGroupResource
	if err := global.Db().Where("id = ? AND group_id = ? AND tenant_id = ?", resourceId, groupId, tenantId).First(&res).Error; err != nil {
		return nil, err
	}
	return &res, nil
}

func CreateResourceGroupResource(tenantId, groupId int64, name string, des string, uid int64) (*model.ResourceGroupResource, error) {
	resource := &model.ResourceGroupResource{
		Id:          uid,
		TenantId:    tenantId,
		GroupId:     groupId,
		Name:        name,
		Description: des,
	}
	if err := global.Db().Create(&resource).Error; err != nil {
		return nil, err
	}
	return resource, nil
}

func UpdateResourceGroupResource(tenantId, groupId, resourceId int64, name, des string) error {
	if err := global.Db().Where("id = ? AND group_id = ? AND tenant_id = ?", resourceId, groupId, tenantId).
		Model(&model.ResourceGroupResource{}).Updates(map[string]any{"name": name, "description": des}).Error; err != nil {
		return err
	}
	return nil
}

func DeleteResourceGroupResource(tenantId, groupId, resourceId int64) error {
	if err := global.Db().Where("id = ? AND group_id = ? AND tenant_id = ?", resourceId, groupId, tenantId).
		Delete(&model.ResourceGroupResource{}).Error; err != nil {
		return err
	}
	return nil
}
