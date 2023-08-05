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
	switch typ {
	case "github":
		return NewGithubIdProvider(clientId, clientSecret, redirectUrl)
	case "qq":
		return NewQqIdProvider(clientId, clientSecret, redirectUrl)
	case "wechat":
		return NewWeChatIdProvider(clientId, clientSecret, redirectUrl)
	case "wecom":
		return NewWeComInternalIdProvider(clientId, clientSecret, redirectUrl)
	case "dingtalk":
		return NewDingTalkIdProvider(clientId, clientSecret, redirectUrl)
	case "lark":
		return NewLarkIdProvider(clientId, clientSecret, redirectUrl)
	case "google":
		return NewGoogleIdProvider(clientId, clientSecret, redirectUrl)
	default:
		return nil
	}
}
