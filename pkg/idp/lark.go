package idp

import (
	"QuickAuth/pkg/utils"
	"time"

	"golang.org/x/oauth2"
)

type LarkIdProvider struct {
	Config *oauth2.Config
}

func NewLarkIdProvider(clientId string, clientSecret string, redirectUrl string) *LarkIdProvider {
	idp := &LarkIdProvider{
		Config: &oauth2.Config{
			Scopes: []string{},
			Endpoint: oauth2.Endpoint{
				TokenURL: "https://open.feishu.cn/open-apis/auth/v3/tenant_access_token/internal",
			},
			ClientID:     clientId,
			ClientSecret: clientSecret,
			RedirectURL:  redirectUrl,
		},
	}

	return idp
}

// GetToken use code get access_token (*operation of getting code ought to be done in front)
func (idp *LarkIdProvider) GetToken(code string) (*oauth2.Token, error) {
	params := &struct {
		AppID     string `json:"app_id"`
		AppSecret string `json:"app_secret"`
	}{idp.Config.ClientID, idp.Config.ClientSecret}
	var appToken struct {
		Code              int    `json:"code"`
		Msg               string `json:"msg"`
		TenantAccessToken string `json:"tenant_access_token"`
		Expire            int    `json:"expire"`
	}
	if err := utils.Post(idp.Config.Endpoint.TokenURL, params, &appToken); err != nil {
		return nil, err
	}

	t := &oauth2.Token{
		AccessToken: appToken.TenantAccessToken,
		TokenType:   "Bearer",
		Expiry:      time.Unix(time.Now().Unix()+int64(appToken.Expire), 0),
	}
	raw := make(map[string]interface{})
	raw["code"] = code
	t = t.WithExtra(raw)

	return t, nil
}

type LarkUserInfo struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		AccessToken      string `json:"access_token"`
		TokenType        string `json:"token_type"`
		ExpiresIn        int    `json:"expires_in"`
		Name             string `json:"name"`
		EnName           string `json:"en_name"`
		AvatarUrl        string `json:"avatar_url"`
		AvatarThumb      string `json:"avatar_thumb"`
		AvatarMiddle     string `json:"avatar_middle"`
		AvatarBig        string `json:"avatar_big"`
		OpenId           string `json:"open_id"`
		UnionId          string `json:"union_id"`
		Email            string `json:"email"`
		UserId           string `json:"user_id"`
		Mobile           string `json:"mobile"`
		TenantKey        string `json:"tenant_key"`
		RefreshExpiresIn int    `json:"refresh_expires_in"`
		RefreshToken     string `json:"refresh_token"`
	} `json:"data"`
}

func (idp *LarkIdProvider) GetUserInfo(token *oauth2.Token) (*UserInfo, error) {
	body := &struct {
		GrantType string `json:"grant_type"`
		Code      string `json:"code"`
	}{"authorization_code", token.Extra("code").(string)}
	var larkUserInfo LarkUserInfo
	if err := utils.Post("https://open.feishu.cn/open-apis/authen/v1/access_token", body, &larkUserInfo,
		map[string]string{"Authorization": "Bearer " + token.AccessToken, "Content-Type": "application/json;charset=UTF-8"}); err != nil {
		return nil, err
	}

	userInfo := UserInfo{
		Id:          larkUserInfo.Data.OpenId,
		DisplayName: larkUserInfo.Data.EnName,
		Username:    larkUserInfo.Data.Name,
		Email:       larkUserInfo.Data.Email,
		AvatarUrl:   larkUserInfo.Data.AvatarUrl,
	}

	return &userInfo, nil
}
