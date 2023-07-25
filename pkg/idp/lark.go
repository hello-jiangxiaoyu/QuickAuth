package idp

import (
	"QuickAuth/pkg/utils"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"time"

	"golang.org/x/oauth2"
)

type LarkIdProvider struct {
	Client *http.Client
	Config *oauth2.Config
}

func NewLarkIdProvider(clientId string, clientSecret string, redirectUrl string) *LarkIdProvider {
	idp := &LarkIdProvider{}

	config := idp.getConfig(clientId, clientSecret, redirectUrl)
	idp.Config = config

	return idp
}

func (idp *LarkIdProvider) SetHttpClient(client *http.Client) {
	idp.Client = client
}

// getConfig return a point of Config, which describes a typical 3-legged OAuth2 flow
func (idp *LarkIdProvider) getConfig(clientId string, clientSecret string, redirectUrl string) *oauth2.Config {
	endpoint := oauth2.Endpoint{
		TokenURL: "https://open.feishu.cn/open-apis/auth/v3/tenant_access_token/internal",
	}

	config := &oauth2.Config{
		Scopes:       []string{},
		Endpoint:     endpoint,
		ClientID:     clientId,
		ClientSecret: clientSecret,
		RedirectURL:  redirectUrl,
	}

	return config
}

type LarkAccessToken struct {
	Code              int    `json:"code"`
	Msg               string `json:"msg"`
	TenantAccessToken string `json:"tenant_access_token"`
	Expire            int    `json:"expire"`
}

// GetToken use code get access_token (*operation of getting code ought to be done in front)
// get more detail via: https://docs.microsoft.com/en-us/linkedIn/shared/authentication/authorization-code-flow?context=linkedIn%2Fcontext&tabs=HTTPS
func (idp *LarkIdProvider) GetToken(code string) (*oauth2.Token, error) {
	params := &struct {
		AppID     string `json:"app_id"`
		AppSecret string `json:"app_secret"`
	}{idp.Config.ClientID, idp.Config.ClientSecret}
	data, err := idp.postWithBody(params, idp.Config.Endpoint.TokenURL)

	appToken := &LarkAccessToken{}
	if err = json.Unmarshal(data, appToken); err != nil || appToken.Code != 0 {
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

// GetUserInfo use LarkAccessToken gotten before return LinkedInUserInfo
// get more detail via: https://docs.microsoft.com/en-us/linkedin/consumer/integrations/self-serve/sign-in-with-linkedin?context=linkedin/consumer/context
func (idp *LarkIdProvider) GetUserInfo(token *oauth2.Token) (*UserInfo, error) {
	body := &struct {
		GrantType string `json:"grant_type"`
		Code      string `json:"code"`
	}{"authorization_code", token.Extra("code").(string)}
	data, _ := json.Marshal(body)
	req, err := http.NewRequest("POST", "https://open.feishu.cn/open-apis/authen/v1/access_token", strings.NewReader(string(data)))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	req.Header.Set("Authorization", "Bearer "+token.AccessToken)

	resp, err := idp.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer utils.DeferErr(resp.Body.Close)
	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var larkUserInfo LarkUserInfo
	if err = json.Unmarshal(data, &larkUserInfo); err != nil {
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

func (idp *LarkIdProvider) postWithBody(body interface{}, url string) ([]byte, error) {
	bs, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	r := strings.NewReader(string(bs))
	resp, err := idp.Client.Post(url, "application/json;charset=UTF-8", r)
	if err != nil {
		return nil, err
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer utils.DeferErr(resp.Body.Close)
	return data, nil
}
