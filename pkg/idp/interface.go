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

func GetIdProvider(typ string, subType string, clientId string, clientSecret string, redirectUrl string) IdProvider {
	if typ == "GitHub" {
		return NewGithubIdProvider(clientId, clientSecret, redirectUrl)
	} else if typ == "QQ" {
		return NewQqIdProvider(clientId, clientSecret, redirectUrl)
	} else if typ == "WeChat" {
		return NewWeChatIdProvider(clientId, clientSecret, redirectUrl)
	} else if typ == "DingTalk" {
		return NewDingTalkIdProvider(clientId, clientSecret, redirectUrl)
	} else if typ == "WeCom" {
		if subType == "Internal" {
			return NewWeComInternalIdProvider(clientId, clientSecret, redirectUrl)
		} else {
			return nil
		}
	} else if typ == "Lark" {
		return NewLarkIdProvider(clientId, clientSecret, redirectUrl)
	}

	return nil
}
