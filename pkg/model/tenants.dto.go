package model

import (
	"time"
)

type TenantDto struct {
	ID         int64     `json:"id"`
	AppID      string    `json:"appId"`
	UserPoolID int64     `json:"userPoolId"`
	Type       int32     `json:"type"`
	Name       string    `json:"name"`
	Host       string    `json:"host"`
	Company    string    `json:"company"`
	Describe   string    `json:"describe"`
	IsDisabled bool      `json:"isDisabled"`
	CreatedAt  time.Time `json:"createdAt"`
}

func (t *Tenant) Dto() *TenantDto {
	return &TenantDto{
		ID:         t.ID,
		AppID:      t.AppID,
		UserPoolID: t.UserPoolID,
		Type:       t.Type,
		Name:       t.Name,
		Host:       t.Host,
		Company:    t.Company,
		Describe:   t.Describe,
		IsDisabled: t.IsDisabled,
		CreatedAt:  t.CreatedAt,
	}
}

func TenantsDto(t Tenant) TenantDto {
	return *t.Dto()
}
