package iam

import (
	"QuickAuth/biz/endpoint/model"
	"QuickAuth/biz/endpoint/request"
	"QuickAuth/pkg/global"
)

func ListResources(tenantId int64) ([]model.Resource, error) {
	var data []model.Resource
	if err := global.Db().Where("tenant_id = ?", tenantId).Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func GetResource(tenantId int64, in *request.Iam) (*model.Resource, error) {
	var data *model.Resource
	if err := global.Db().Where("tenant_id = ?", tenantId, in.ResourceId).First(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func CreateResource(res *model.Resource) (*model.Resource, error) {
	if err := global.Db().Create(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}

func UpdateResource(tenantId int64, resId int64, res *model.Resource) error {
	if err := global.Db().Where("tenant_id = ? AND id = ?", tenantId, resId).Updates(res).Error; err != nil {
		return err
	}

	return nil
}

func DeleteResource(tenantId int64, resId int64) error {
	if err := global.Db().Where("tenant_id = ? AND id = ?", tenantId, resId).Delete(model.Resource{}).Error; err != nil {
		return err
	}

	return nil
}
