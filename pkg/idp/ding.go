package idp

import (
	"QuickAuth/pkg/utils"
	"fmt"
	"time"

	"golang.org/x/oauth2"
)

type DingTalkIdProvider struct {
	Config *oauth2.Config
}

// NewDingTalkIdProvider ...
func NewDingTalkIdProvider(clientId string, clientSecret string, redirectUrl string) *DingTalkIdProvider {
	idp := &DingTalkIdProvider{
		Config: &oauth2.Config{
			Scopes: []string{"", ""}, // DingTalk not allow to set scopes,here it is just a placeholder,
			Endpoint: oauth2.Endpoint{
				AuthURL:  "https://api.dingtalk.com/v1.0/contact/users/me",
				TokenURL: "https://api.dingtalk.com/v1.0/oauth2/userAccessToken",
			},
			ClientID:     clientId,
			ClientSecret: clientSecret,
			RedirectURL:  redirectUrl,
		},
	}

	return idp
}

// GetToken use code get access_token (*operation of getting authCode ought to be done in front)
func (idp *DingTalkIdProvider) GetToken(code string) (*oauth2.Token, error) {
	pTokenParams := &struct {
		ClientId     string `json:"clientId"`
		ClientSecret string `json:"clientSecret"`
		Code         string `json:"code"`
		GrantType    string `json:"grantType"`
	}{idp.Config.ClientID, idp.Config.ClientSecret, code, "authorization_code"}
	var pToken struct {
		ErrCode     int    `json:"code"`
		ErrMsg      string `json:"message"`
		AccessToken string `json:"accessToken"` // Interface call credentials
		ExpiresIn   int64  `json:"expireIn"`    // access_token interface call credential timeout time, unit (seconds)
	}
	if err := utils.Post(idp.Config.Endpoint.TokenURL, pTokenParams, &pToken); err != nil {
		return nil, err
	}

	if pToken.ErrCode != 0 {
		return nil, fmt.Errorf("pToken.Errcode = %d, pToken.Errmsg = %s", pToken.ErrCode, pToken.ErrMsg)
	}

	token := &oauth2.Token{
		AccessToken: pToken.AccessToken,
		Expiry:      time.Unix(time.Now().Unix()+pToken.ExpiresIn, 0),
	}
	return token, nil
}

// GetUserInfo Use access_token to get UserInfo
func (idp *DingTalkIdProvider) GetUserInfo(token *oauth2.Token) (*UserInfo, error) {
	var dtUserInfo struct {
		Nick      string `json:"nick"`
		OpenId    string `json:"openId"`
		UnionId   string `json:"unionId"`
		AvatarUrl string `json:"avatarUrl"`
		Email     string `json:"email"`
		Mobile    string `json:"mobile"`
		StateCode string `json:"stateCode"`
	}
	if err := utils.Get(idp.Config.Endpoint.AuthURL, &dtUserInfo,
		map[string]string{"x-acs-dingtalk-access-token": token.AccessToken}); err != nil {
		return nil, err
	}
	countryCode, err := utils.GetCountryCode(dtUserInfo.StateCode, dtUserInfo.Mobile)
	if err != nil {
		return nil, err
	}

	userInfo := UserInfo{
		Id:          dtUserInfo.OpenId,
		Username:    dtUserInfo.Nick,
		DisplayName: dtUserInfo.Nick,
		UnionId:     dtUserInfo.UnionId,
		Email:       dtUserInfo.Email,
		Phone:       dtUserInfo.Mobile,
		CountryCode: countryCode,
		AvatarUrl:   dtUserInfo.AvatarUrl,
	}
	corpAccessToken, err := idp.getInnerAppAccessToken()
	if err != nil {
		return nil, err
	}
	userId, err := idp.getUserId(userInfo.UnionId, corpAccessToken)
	if err != nil {
		return nil, err
	}

	if corpMobile, corpEmail, jobNumber, err := idp.getUserCorpEmail(userId, corpAccessToken); err == nil {
		if corpMobile != "" {
			userInfo.Phone = corpMobile
		}
		if corpEmail != "" {
			userInfo.Email = corpEmail
		}
		if jobNumber != "" {
			userInfo.Username = jobNumber
		}
	}
	return &userInfo, nil
}

func (idp *DingTalkIdProvider) getInnerAppAccessToken() (string, error) {
	body := make(map[string]string)
	body["appKey"] = idp.Config.ClientID
	body["appSecret"] = idp.Config.ClientSecret
	var data struct {
		ExpireIn    int    `json:"expireIn"`
		AccessToken string `json:"accessToken"`
	}
	if err := utils.Post("https://api.dingtalk.com/v1.0/oauth2/accessToken", body, &data); err != nil {
		return "", err
	}

	return data.AccessToken, nil
}

func (idp *DingTalkIdProvider) getUserId(unionId string, accessToken string) (string, error) {
	body := make(map[string]string)
	body["unionid"] = unionId
	var data struct {
		ErrCode    int    `json:"errcode"`
		ErrMessage string `json:"errmsg"`
		Result     struct {
			UserId string `json:"userid"`
		} `json:"result"`
	}
	if err := utils.Post("https://oapi.dingtalk.com/topapi/user/getbyunionid?access_token="+accessToken, body, &data); err != nil {
		return "", err
	}

	if data.ErrCode == 60121 {
		return "", fmt.Errorf("该应用只允许本企业内部用户登录，您不属于该企业，无法登录")
	} else if data.ErrCode != 0 {
		return "", fmt.Errorf(data.ErrMessage)
	}
	return data.Result.UserId, nil
}

func (idp *DingTalkIdProvider) getUserCorpEmail(userId string, accessToken string) (string, string, string, error) {
	body := make(map[string]string)
	body["userid"] = userId
	var data struct {
		ErrMessage string `json:"errmsg"`
		Result     struct {
			Mobile    string `json:"mobile"`
			Email     string `json:"email"`
			JobNumber string `json:"job_number"`
		} `json:"result"`
	}
	if err := utils.Post("https://oapi.dingtalk.com/topapi/v2/user/get?access_token="+accessToken, body, &data); err != nil {
		return "", "", "", err
	}
	if data.ErrMessage != "ok" {
		return "", "", "", fmt.Errorf(data.ErrMessage)
	}

	return data.Result.Mobile, data.Result.Email, data.Result.JobNumber, nil
}
