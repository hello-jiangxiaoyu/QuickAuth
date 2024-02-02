package idp

import (
	"QuickAuth/pkg/utils"
	"context"
	"errors"
	"fmt"
	"golang.org/x/oauth2"
)

type GoogleIdProvider struct {
	Config *oauth2.Config
}

func NewGoogleIdProvider(clientId string, clientSecret string, redirectUrl string) *GoogleIdProvider {
	idp := &GoogleIdProvider{
		Config: &oauth2.Config{
			Scopes: []string{"profile", "email"},
			Endpoint: oauth2.Endpoint{
				AuthURL:  "https://accounts.google.com/o/oauth2/auth",
				TokenURL: "https://accounts.google.com/o/oauth2/token",
			},
			ClientID:     clientId,
			ClientSecret: clientSecret,
			RedirectURL:  redirectUrl,
		},
	}

	return idp
}

func (idp *GoogleIdProvider) GetToken(code string) (*oauth2.Token, error) {
	ctx := context.WithValue(context.Background(), oauth2.HTTPClient, utils.DefaultClient)
	return idp.Config.Exchange(ctx, code)
}

func (idp *GoogleIdProvider) GetUserInfo(token *oauth2.Token) (*UserInfo, error) {
	var googleUserInfo struct {
		Id            string `json:"id"`
		Email         string `json:"email"`
		VerifiedEmail bool   `json:"verified_email"`
		Name          string `json:"name"`
		GivenName     string `json:"given_name"`
		FamilyName    string `json:"family_name"`
		Picture       string `json:"picture"`
		Locale        string `json:"locale"`
	}
	url := fmt.Sprintf("https://www.googleapis.com/oauth2/v2/userinfo?alt=json&access_token=%s", token.AccessToken)
	if err := utils.Get(url, &googleUserInfo); err != nil {
		return nil, err
	}
	if googleUserInfo.Email == "" {
		return nil, errors.New("google email is empty")
	}

	userInfo := UserInfo{
		Id:          googleUserInfo.Id,
		Username:    googleUserInfo.Email,
		DisplayName: googleUserInfo.Name,
		Email:       googleUserInfo.Email,
		AvatarUrl:   googleUserInfo.Picture,
	}
	return &userInfo, nil
}
