package service

import (
	"QuickAuth/internal/service/admin"
	"QuickAuth/internal/service/oauth"
	"QuickAuth/pkg/global"
)

type Service struct {
	oauth.ServiceOauth
	admin.ServiceAdmin
}

func NewService(repo *global.Repository) *Service {
	a := admin.NewAdminService(repo)
	o := oauth.NewOauthService(repo)
	return &Service{
		ServiceAdmin: *a,
		ServiceOauth: *o,
	}
}
