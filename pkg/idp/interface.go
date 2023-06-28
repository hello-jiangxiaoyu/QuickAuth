package idp

import (
	"golang.org/x/oauth2"
	"net/http"
)

type UserInfo struct {
	Id          string
	Username    string
	DisplayName string
	UnionId     string
	Email       string
	Phone       string
	CountryCode string
	AvatarUrl   string
}

type IdProvider interface {
	SetHttpClient(client *http.Client)
	GetToken(code string) (*oauth2.Token, error)
	GetUserInfo(token *oauth2.Token) (*UserInfo, error)
}

func GetIdProvider(typ string, clientId string, clientSecret string, redirectUrl string) IdProvider {
	if typ == "github" {
		return NewGithubIdProvider(clientId, clientSecret, redirectUrl)
	} else if typ == "qq" {
		return NewQqIdProvider(clientId, clientSecret, redirectUrl)
	} else if typ == "wechat" {
		return NewWeChatIdProvider(clientId, clientSecret, redirectUrl)
	} else if typ == "dingtalk" {
		return NewDingTalkIdProvider(clientId, clientSecret, redirectUrl)
	} else if typ == "wecom" {
		return NewWeComInternalIdProvider(clientId, clientSecret, redirectUrl)
	} else if typ == "lark" {
		return NewLarkIdProvider(clientId, clientSecret, redirectUrl)
	}

	return nil
}
