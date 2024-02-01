package iam

import "QuickAuth/biz/endpoint/model"

func (s *ServiceIam) ListResourceNodes(tenantId int64, resId int64) ([]model.ResourceNode, error) {
	var data []model.ResourceNode
	if err := s.db.Where("tenant_id = ? AND resource_id = ?", tenantId, resId).
		Find(&data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

func (s *ServiceIam) GetResourceNode(tenantId int64, resId int64, nodeId int64) (*model.ResourceNode, error) {
	var data model.ResourceNode
	if err := s.db.Where("id = ? AND resource_id = ? AND tenant_id = ?", nodeId, resId, tenantId).
		First(&data).Error; err != nil {
		return nil, err
	}

	return &data, nil
}

func (s *ServiceIam) CreateResourceNode(node *model.ResourceNode) (*model.ResourceNode, error) {
	if err := s.db.Create(node).Error; err != nil {
		return nil, err
	}

	return node, nil
}

func (s *ServiceIam) UpdateResourceNode(tenantId int64, resId int64, nodeId int64, node *model.ResourceNode) error {
	if err := s.db.Where("id = ? AND resId = ? AND tenant_id = ?", nodeId, resId, tenantId).
		Updates(node).Error; err != nil {
		return err
	}

	return nil
}

func (s *ServiceIam) DeleteResourceNode(tenantId int64, resId int64, nodeId int64) error {
	if err := s.db.Where("id = ? AND resId = ? AND tenant_id = ?", nodeId, resId, tenantId).
		Delete(model.ResourceNode{}).Error; err != nil {
		return err
	}

	return nil
}
