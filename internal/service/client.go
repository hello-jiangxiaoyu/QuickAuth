package service

import (
	"QuickAuth/pkg/model"
	"QuickAuth/pkg/tools/utils"
)

func (s *Service) ListClients() ([]model.Client, error) {
	var clients []model.Client
	if err := s.db.Select("id", "name").Find(&clients).Error; err != nil {
		return nil, err
	}

	return clients, nil
}

func (s *Service) GetClient(id string) (*model.Client, error) {
	var client model.Client
	if err := s.db.Where("id = ?", id).First(&client).Error; err != nil {
		return nil, err
	}
	return &client, nil
}

func (s *Service) CreateClient(client model.Client) (*model.Client, error) {
	client.ID = utils.GetNoLineUUID()
	if err := s.db.Create(&client).Error; err != nil {
		return nil, err
	}
	return &client, nil
}

func (s *Service) ModifyClient(client model.Client) error {
	if err := s.db.Where("id = ?", client.ID).Save(client).Error; err != nil {
		return err
	}
	return nil
}

func (s *Service) DeleteClient(clientId string) error {
	if _, err := s.GetClient(clientId); err != nil {
		return err
	}
	if err := s.db.Where("id = ?", clientId).Delete(model.Client{}).Error; err != nil {
		return err
	}
	return nil
}

// =================== client secret ===================

func (s *Service) ListClientSecrets(clientId string) ([]model.ClientSecret, error) {
	var secrets []model.ClientSecret
	if err := s.db.Where("client_id = ?", clientId).Find(&secrets).Error; err != nil {
		return nil, err
	}

	return secrets, nil
}

func (s *Service) CreateClientSecret(secret model.ClientSecret) (*model.ClientSecret, error) {
	if err := s.db.Create(&secret).Error; err != nil {
		return nil, err
	}
	return &secret, nil
}

func (s *Service) ModifyClientSecret(secret model.ClientSecret) error {
	if err := s.db.Where("id = ? AND client_id = ?", secret.ID, secret.ClientID).Save(secret).Error; err != nil {
		return err
	}
	return nil
}

func (s *Service) DeleteClientSecret(clientId, secretId string) error {
	if err := s.db.Where("id = ? AND client_id = ?", clientId, secretId).
		Delete(&model.ClientSecret{}).Error; err != nil {
		return err
	}
	return nil
}

// =================== redirect uri ===================

func (s *Service) IsRedirectUriValid(clientId, uri string) (bool, error) {
	var redirectUri model.RedirectURI
	if err := s.db.Select("uri").Where("client_id = ? AND uri = ?", clientId, uri).
		Limit(1).Find(&uri).Error; err != nil {
		return false, err
	}

	return redirectUri.URI == uri, nil
}

func (s *Service) ListRedirectUri(clientId string) ([]model.RedirectURI, error) {
	var clients []model.RedirectURI
	if err := s.db.Where("client_id = ?", clientId).Find(&clients).Error; err != nil {
		return nil, err
	}

	return clients, nil
}

func (s *Service) CreateRedirectUri(uri model.RedirectURI) (*model.RedirectURI, error) {
	if err := s.db.Create(&uri).Error; err != nil {
		return nil, err
	}
	return &uri, nil
}

func (s *Service) ModifyRedirectUri(uri model.RedirectURI) error {
	if err := s.db.Where("id = ? AND client_id = ?", uri.ID, uri.ClientID).Save(uri).Error; err != nil {
		return err
	}
	return nil
}

func (s *Service) DeleteRedirectUri(clientId, uriId string) error {
	if err := s.db.Where("id = ? AND client_id = ?", clientId, uriId).
		Delete(&model.RedirectURI{}).Error; err != nil {
		return err
	}
	return nil
}
