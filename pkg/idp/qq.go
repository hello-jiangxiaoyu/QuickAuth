package idp

import (
	"QuickAuth/pkg/utils"
	"errors"
	"fmt"
	"io"
	"net/url"
	"regexp"

	"golang.org/x/oauth2"
)

type QqIdProvider struct {
	Config *oauth2.Config
}

func NewQqIdProvider(clientId string, clientSecret string, redirectUrl string) *QqIdProvider {
	idp := &QqIdProvider{
		Config: &oauth2.Config{
			Scopes: []string{"get_user_info"},
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

var re = regexp.MustCompile("token=(.*?)&")

func (idp *QqIdProvider) GetToken(code string) (*oauth2.Token, error) {
	params := url.Values{}
	params.Add("grant_type", "authorization_code")
	params.Add("client_id", idp.Config.ClientID)
	params.Add("client_secret", idp.Config.ClientSecret)
	params.Add("code", code)
	params.Add("redirect_uri", idp.Config.RedirectURL)
	accessTokenUrl := fmt.Sprintf("https://graph.qq.com/oauth2.0/token?%s", params.Encode())
	resp, err := utils.DefaultClient.Get(accessTokenUrl)
	if err != nil {
		return nil, err
	}

	defer utils.DeferErr(resp.Body.Close)
	tokenContent, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	matched := re.FindAllStringSubmatch(string(tokenContent), -1)
	accessToken := matched[0][1]
	token := &oauth2.Token{
		AccessToken: accessToken,
		TokenType:   "Bearer",
	}
	return token, nil
}

type QqUserInfo struct {
	Ret             int    `json:"ret"`
	Msg             string `json:"msg"`
	IsLost          int    `json:"is_lost"`
	Nickname        string `json:"nickname"`
	Gender          string `json:"gender"`
	GenderType      int    `json:"gender_type"`
	Province        string `json:"province"`
	City            string `json:"city"`
	Year            string `json:"year"`
	Constellation   string `json:"constellation"`
	Figureurl       string `json:"figureurl"`
	Figureurl1      string `json:"figureurl_1"`
	Figureurl2      string `json:"figureurl_2"`
	FigureurlQq1    string `json:"figureurl_qq_1"`
	FigureurlQq2    string `json:"figureurl_qq_2"`
	FigureurlQq     string `json:"figureurl_qq"`
	FigureurlType   string `json:"figureurl_type"`
	IsYellowVip     string `json:"is_yellow_vip"`
	Vip             string `json:"vip"`
	YellowVipLevel  string `json:"yellow_vip_level"`
	Level           string `json:"level"`
	IsYellowYearVip string `json:"is_yellow_year_vip"`
}

func (idp *QqIdProvider) GetUserInfo(token *oauth2.Token) (*UserInfo, error) {
	openIdUrl := fmt.Sprintf("https://graph.qq.com/oauth2.0/me?access_token=%s", token.AccessToken)
	resp, err := utils.DefaultClient.Get(openIdUrl)
	if err != nil {
		return nil, err
	}
	defer utils.DeferErr(resp.Body.Close)
	openIdBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	re2 := regexp.MustCompile("\"openid\":\"(.*?)\"}")
	matched := re2.FindAllStringSubmatch(string(openIdBody), -1)
	openId := matched[0][1]
	if openId == "" {
		return nil, errors.New("openId is empty")
	}

	var qqUserInfo QqUserInfo
	userInfoUrl := fmt.Sprintf("https://graph.qq.com/user/get_user_info?access_token=%s&oauth_consumer_key=%s&openid=%s", token.AccessToken, idp.Config.ClientID, openId)
	if err = utils.Get(userInfoUrl, &qqUserInfo); err != nil {
		return nil, err
	}

	if qqUserInfo.Ret != 0 {
		return nil, fmt.Errorf("ret expected 0, got %d", qqUserInfo.Ret)
	}

	userInfo := UserInfo{
		Id:          openId,
		Username:    qqUserInfo.Nickname,
		DisplayName: qqUserInfo.Nickname,
		AvatarUrl:   qqUserInfo.FigureurlQq1,
	}
	return &userInfo, nil
}
