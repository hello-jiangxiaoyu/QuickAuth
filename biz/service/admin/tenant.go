package admin

import (
	"QuickAuth/biz/endpoint/model"
	"QuickAuth/pkg/global"
	"QuickAuth/pkg/utils"
	"github.com/lib/pq"
	"net/url"
)

func GetTenant(_ string, tenantId int64) (*model.Tenant, error) {
	var tenant model.Tenant
	if err := global.Db().Where("id = ?", tenantId).Preload("App").Preload("UserPool").First(&tenant).Error; err != nil {
		return nil, err
	}
	return &tenant, nil
}

func ListTenant(appId string) ([]model.Tenant, error) {
	var tenant []model.Tenant
	if err := global.Db().Select("id", "app_id", "user_pool_id", "type", "host", "name", "company", "created_at").
		Where("app_id = ?", appId).Find(&tenant).Error; err != nil {
		return nil, err
	}
	return tenant, nil
}

func CreatTenant(t *model.Tenant) (*model.Tenant, error) {
	if _, err := GetApp(t.AppID); err != nil {
		return nil, utils.WithMessage(err, "no such app")
	}
	if _, err := GetUserPool(t.UserPoolID); err != nil {
		return nil, utils.WithMessage(err, "no such user pool")
	}

	t.RedirectUris = pq.StringArray{"https://" + t.Host, "http://" + t.Host}
	t.GrantTypes = pq.StringArray{""}
	if err := global.Db().Select("app_id", "user_pool_id", "type", "name", "host", "company", "grant_types", "redirect_uris", "describe").
		Create(t).Error; err != nil {
		return nil, err
	}
	return t, nil
}

func ModifyTenant(tenantId int64, t *model.Tenant) error {
	if err := global.Db().Select("type", "name", "host", "company", "grant_type", "describe",
		"is_code", "is_refresh", "is_password", "is_credential", "is_device_flow",
		"code_expire", "id_expire", "access_expire", "refresh_expire").
		Where("id = ? AND app_id = ?", tenantId, t.AppID).Updates(t).Error; err != nil {
		return err
	}
	return nil
}

func DeleteTenant(appId string, tenantId int64) error {
	if err := global.Db().Where("id = ? AND app_id = ?", tenantId, appId).Delete(&model.Tenant{}).Error; err != nil {
		return err
	}
	return nil
}

// =================== redirect uri ===================

func IsRedirectUriValid(appId string, tenantId int64, uri string) (bool, error) {
	var app model.Tenant
	if err := global.Db().Select("uri").Where("id = ? AND app_id = ?", tenantId, appId).First(&app).Error; err != nil {
		return false, err
	}

	for _, v := range app.RedirectUris {
		if v == uri {
			return true, nil
		}
	}
	return false, nil
}

func ListRedirectUri(tenantId int64) ([]string, error) {
	var apps []string
	if err := global.Db().Model(model.Tenant{}).Select("redirect_uris").
		Where("id = ?", tenantId).Find(&apps).Error; err != nil {
		return nil, err
	}

	return apps, nil
}

func CreateRedirectUri(tenantId int64, uri string) error {
	sql := `update tenants set redirect_uris = array_prepend(?, redirect_uris) where id = ?;`
	if err := global.Db().Exec(sql, uri, tenantId).Error; err != nil {
		return err
	}

	return nil
}

func ModifyRedirectUri(tenantId int64, uriId uint, uri string) error {
	sql := `update tenants set redirect_uris[?] = ? where id = ?;`
	if err := global.Db().Exec(sql, uriId, uri, tenantId).Error; err != nil {
		return err
	}

	return nil
}

func DeleteRedirectUri(tenantId int64, uri string) error {
	uri, err := url.QueryUnescape(uri)
	if err != nil {
		return utils.WithMessage(err, "invalid uri")
	}

	sql := `update tenants set redirect_uris = array_remove(redirect_uris, ?) where id = ?;`
	if err = global.Db().Exec(sql, uri, tenantId).Error; err != nil {
		return err
	}

	return nil
}
