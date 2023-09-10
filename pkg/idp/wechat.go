package idp

import (
	"QuickAuth/pkg/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"time"

	"golang.org/x/oauth2"
)

type WeChatIdProvider struct {
	Config *oauth2.Config
}

func NewWeChatIdProvider(clientId string, clientSecret string, redirectUrl string) *WeChatIdProvider {
	idp := &WeChatIdProvider{
		Config: &oauth2.Config{
			Scopes: []string{"snsapi_login"},
			Endpoint: oauth2.Endpoint{
				TokenURL: "https://graph.qq.com/oauth2.0/token",
			},
			ClientID:     clientId,
			ClientSecret: clientSecret,
			RedirectURL:  redirectUrl,
		},
	}

	return idp
}

type WechatAccessToken struct {
	AccessToken  string `json:"access_token"`  // Interface call credentials
	ExpiresIn    int64  `json:"expires_in"`    // access_token interface call credential timeout time, unit (seconds)
	RefreshToken string `json:"refresh_token"` // User refresh access_token
	Openid       string `json:"openid"`        // Unique ID of authorized user
	Scope        string `json:"scope"`         // The scope of user authorization, separated by commas. (,)
	Unionid      string `json:"unionid"`       // This field will appear if and only if the website application has been authorized by the user's UserInfo.
}

// GetToken use code get access_token (*operation of getting code ought to be done in front)
func (idp *WeChatIdProvider) GetToken(code string) (*oauth2.Token, error) {
	params := url.Values{}
	params.Add("grant_type", "authorization_code")
	params.Add("appid", idp.Config.ClientID)
	params.Add("secret", idp.Config.ClientSecret)
	params.Add("code", code)
	tokenResponse, err := utils.DefaultClient.Get(fmt.Sprintf("https://api.weixin.qq.com/sns/oauth2/access_token?%s", params.Encode()))
	if err != nil {
		return nil, err
	}

	defer utils.DeferErr(tokenResponse.Body.Close)
	buf := new(bytes.Buffer)
	if _, err = buf.ReadFrom(tokenResponse.Body); err != nil {
		return nil, err
	}
	if strings.Contains(buf.String(), "errcode") {
		return nil, fmt.Errorf(buf.String())
	}

	var wechatAccessToken WechatAccessToken
	if err = json.Unmarshal(buf.Bytes(), &wechatAccessToken); err != nil {
		return nil, err
	}

	token := oauth2.Token{
		AccessToken:  wechatAccessToken.AccessToken,
		TokenType:    "WeChatAccessToken",
		RefreshToken: wechatAccessToken.RefreshToken,
		Expiry:       time.Time{},
	}
	raw := make(map[string]string)
	raw["Openid"] = wechatAccessToken.Openid
	token.WithExtra(raw)

	return &token, nil
}

type WechatUserInfo struct {
	Openid     string   `json:"openid"`   // The ID of an ordinary user, which is unique to the current developer account
	Nickname   string   `json:"nickname"` // Ordinary user nickname
	Sex        int      `json:"sex"`      // Ordinary user gender, 1 is male, 2 is female
	Language   string   `json:"language"`
	City       string   `json:"city"`       // City filled in by general user's personal data
	Province   string   `json:"province"`   // Province filled in by ordinary user's personal information
	Country    string   `json:"country"`    // Country, such as China is CN
	Headimgurl string   `json:"headimgurl"` // User avatar, the last value represents the size of the square avatar (there are optional values of 0, 46, 64, 96, 132, 0 represents a 640*640 square avatar), this item is empty when the user does not have an avatar
	Privilege  []string `json:"privilege"`  // User Privilege information, json array, such as Wechat Woka user (chinaunicom)
	Unionid    string   `json:"unionid"`    // Unified user identification. For an application under a WeChat open platform account, the unionid of the same user is unique.
}

// GetUserInfo use WechatAccessToken gotten before return WechatUserInfo
func (idp *WeChatIdProvider) GetUserInfo(token *oauth2.Token) (*UserInfo, error) {
	var wechatUserInfo WechatUserInfo
	accessToken := token.AccessToken
	openid := token.Extra("Openid")
	resp, err := utils.DefaultClient.Get(fmt.Sprintf("https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s", accessToken, openid))
	if err != nil {
		return nil, err
	}

	defer utils.DeferErr(resp.Body.Close)
	buf := new(bytes.Buffer)
	if _, err = buf.ReadFrom(resp.Body); err != nil {
		return nil, err
	}
	if err = json.Unmarshal(buf.Bytes(), &wechatUserInfo); err != nil {
		return nil, err
	}

	id := wechatUserInfo.Unionid
	if id == "" {
		id = wechatUserInfo.Openid
	}
	userInfo := UserInfo{
		Id:          id,
		Username:    wechatUserInfo.Nickname,
		DisplayName: wechatUserInfo.Nickname,
		AvatarUrl:   wechatUserInfo.Headimgurl,
	}
	return &userInfo, nil
}
