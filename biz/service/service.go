package service

import (
	"QuickAuth/biz/service/admin"
	"QuickAuth/biz/service/iam"
	"QuickAuth/biz/service/oauth"
	"QuickAuth/pkg/global"
)

type Service struct {
	oauth.ServiceOauth
	admin.ServiceAdmin
	iam.ServiceIam
}

func NewService(repo *global.Repository) *Service {
	a := admin.NewAdminService(repo)
	o := oauth.NewOauthService(repo)
	i := iam.NewIamService(repo)
	return &Service{
		ServiceAdmin: *a,
		ServiceOauth: *o,
		ServiceIam:   *i,
	}
}
