package idp

import (
	"QuickAuth/pkg/utils"
	"context"
	"fmt"
	_ "net/url"
	_ "time"

	"golang.org/x/oauth2"
)

type CustomIdProvider struct {
	Config      *oauth2.Config
	UserInfoUrl string
}

func NewAuth2IdProvider(clientId string, clientSecret string, redirectUrl string, authUrl string, tokenUrl string, userInfoUrl string) *CustomIdProvider {
	idp := &CustomIdProvider{
		Config: &oauth2.Config{
			Endpoint: oauth2.Endpoint{
				AuthURL:  authUrl,
				TokenURL: tokenUrl,
			},
			ClientID:     clientId,
			ClientSecret: clientSecret,
			RedirectURL:  redirectUrl,
		},
		UserInfoUrl: userInfoUrl,
	}

	return idp
}

func (idp *CustomIdProvider) GetToken(code string) (*oauth2.Token, error) {
	ctx := context.WithValue(context.Background(), oauth2.HTTPClient, utils.DefaultClient)
	return idp.Config.Exchange(ctx, code)
}

type CustomUserInfo struct {
	Id          string `json:"sub"`
	Name        string `json:"preferred_username,omitempty"`
	DisplayName string `json:"name"`
	Email       string `json:"email"`
	AvatarUrl   string `json:"picture"`
	Status      string `json:"status"`
	Msg         string `json:"msg"`
}

func (idp *CustomIdProvider) GetUserInfo(token *oauth2.Token) (*UserInfo, error) {
	var ctUserinfo struct {
		Id          string `json:"sub"`
		Name        string `json:"preferred_username,omitempty"`
		DisplayName string `json:"name"`
		Email       string `json:"email"`
		AvatarUrl   string `json:"picture"`
		Status      string `json:"status"`
		Msg         string `json:"msg"`
	}

	if err := utils.Get(idp.UserInfoUrl, &ctUserinfo,
		map[string]string{"Authorization": fmt.Sprintf("Bearer %s", token.AccessToken)}); err != nil {
		return nil, err
	}
	if ctUserinfo.Status != "" {
		return nil, fmt.Errorf("err: %s", ctUserinfo.Msg)
	}

	userInfo := &UserInfo{
		Id:          ctUserinfo.Id,
		Username:    ctUserinfo.Name,
		DisplayName: ctUserinfo.DisplayName,
		Email:       ctUserinfo.Email,
		AvatarUrl:   ctUserinfo.AvatarUrl,
	}
	return userInfo, nil
}
