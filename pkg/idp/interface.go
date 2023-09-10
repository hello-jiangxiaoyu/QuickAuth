package idp

import (
	"QuickAuth/internal/endpoint/model"
	"golang.org/x/oauth2"
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
	GetToken(code string) (*oauth2.Token, error)
	GetUserInfo(token *oauth2.Token) (*UserInfo, error)
}

func GetIdProvider(prd *model.Provider, redirectUrl string) IdProvider {
	switch prd.Type {
	case "github":
		return NewGithubIdProvider(prd.ClientID, prd.ClientSecret, redirectUrl)
	case "qq":
		return NewQqIdProvider(prd.ClientID, prd.ClientSecret, redirectUrl)
	case "wechat": // 微信
		return NewWeChatIdProvider(prd.ClientID, prd.ClientSecret, redirectUrl)
	case "wecom": // 企业微信
		return NewWeComInternalIdProvider(prd.ClientID, prd.ClientSecret, redirectUrl)
	case "dingtalk": // 钉钉
		return NewDingTalkIdProvider(prd.ClientID, prd.ClientSecret, redirectUrl)
	case "lark": // 飞书
		return NewLarkIdProvider(prd.ClientID, prd.ClientSecret, redirectUrl)
	case "google": // 谷歌账号
		return NewGoogleIdProvider(prd.ClientID, prd.ClientSecret, redirectUrl)
	case "oauth2": // 支持oauth2协议的其他平台
		return NewAuth2IdProvider(prd.ClientID, prd.ClientSecret, redirectUrl, prd.AuthEndpoint, prd.TokenEndpoint, prd.UserInfoEndpoint)
	default:
		return nil
	}
}
