package service

import (
	"QuickAuth/pkg/model"
	"QuickAuth/pkg/tools/safe"
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
	if _, err := s.GetClient(secret.ClientID); err != nil {
		return nil, err
	}
	secret.Secret = safe.Rand62(31)
	if err := s.db.Create(&secret).Error; err != nil {
		return nil, err
	}
	return &secret, nil
}

func (s *Service) DeleteClientSecret(clientId, secretId string) error {
	if err := s.db.Where("id = ? AND client_id = ?", secretId, clientId).
		Delete(&model.ClientSecret{}).Error; err != nil {
		return err
	}
	return nil
}

// =================== redirect uri ===================

func (s *Service) IsRedirectUriValid(clientId, uri string) (bool, error) {
	var client model.Client
	if err := s.db.Select("uri").Where("id = ?", clientId, uri).First(&client).Error; err != nil {
		return false, err
	}

	for _, v := range client.RedirectUris {
		if v == uri {
			return true, nil
		}
	}
	return false, nil
}

func (s *Service) ListRedirectUri(clientId string) ([]string, error) {
	var clients []string
	if err := s.db.Model(model.Client{}).Select("redirect_uris").Where("id = ?", clientId).Find(&clients).Error; err != nil {
		return nil, err
	}

	return clients, nil
}

func (s *Service) CreateRedirectUri(clientId, uri string) error {
	sql := `update clients set redirect_uris = array_prepend(?, redirect_uris) where id = ?;`
	if err := s.db.Exec(sql, uri, clientId).Error; err != nil {
		return err
	}

	return nil
}

func (s *Service) ModifyRedirectUri(clientId string, uriId uint, uri string) error {
	sql := `update clients set redirect_uris[?] = ? where id = ?;`
	if err := s.db.Exec(sql, uriId, uri, clientId).Error; err != nil {
		return err
	}

	return nil
}

func (s *Service) DeleteRedirectUri(clientId string, uri string) error {
	sql := `update clients set redirect_uris = array_remove(redirect_uris, ?) where id = ?;`
	if err := s.db.Exec(sql, uri, clientId).Error; err != nil {
		return err
	}

	return nil
}
