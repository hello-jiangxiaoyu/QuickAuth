package model

import "QuickAuth/pkg/utils"

type AppDto struct {
	ID       string      `json:"id"`
	Tag      string      `json:"tag"`
	Name     string      `json:"name"`
	Describe string      `json:"describe"`
	Icon     string      `json:"icon"`
	Tenant   []TenantDto `json:"tenant"`
}

func (a *App) Dto() *AppDto {
	return &AppDto{
		ID:       a.ID,
		Tag:      a.Tag,
		Name:     a.Name,
		Describe: a.Describe,
		Icon:     a.Icon,
		Tenant:   utils.DtoFilter(a.Tenant, TenantsDto),
	}
}
