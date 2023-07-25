package idp

import (
	"QuickAuth/pkg/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"golang.org/x/oauth2"
)

type GithubIdProvider struct {
	Client *http.Client
	Config *oauth2.Config
}

func NewGithubIdProvider(clientId string, clientSecret string, redirectUrl string) *GithubIdProvider {
	idp := &GithubIdProvider{}

	config := idp.getConfig()
	config.ClientID = clientId
	config.ClientSecret = clientSecret
	config.RedirectURL = redirectUrl
	idp.Config = config

	return idp
}

func (idp *GithubIdProvider) SetHttpClient(client *http.Client) {
	idp.Client = client
}

func (idp *GithubIdProvider) getConfig() *oauth2.Config {
	endpoint := oauth2.Endpoint{
		AuthURL:  "https://github.com/login/oauth/authorize",
		TokenURL: "https://github.com/login/oauth/access_token",
	}

	config := &oauth2.Config{
		Scopes:   []string{"user:email", "read:user"},
		Endpoint: endpoint,
	}

	return config
}

type GithubToken struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
	Error       string `json:"error"`
}

func (idp *GithubIdProvider) GetToken(code string) (*oauth2.Token, error) {
	params := &struct {
		Code         string `json:"code"`
		ClientId     string `json:"client_id"`
		ClientSecret string `json:"client_secret"`
	}{code, idp.Config.ClientID, idp.Config.ClientSecret}
	data, err := idp.postWithBody(params, idp.Config.Endpoint.TokenURL)
	if err != nil {
		return nil, err
	}
	pToken := &GithubToken{}
	if err = json.Unmarshal(data, pToken); err != nil {
		return nil, err
	}
	if pToken.Error != "" {
		return nil, fmt.Errorf("err: %s", pToken.Error)
	}

	token := &oauth2.Token{
		AccessToken: pToken.AccessToken,
		TokenType:   "Bearer",
	}

	return token, nil
}

type GitHubUserInfo struct {
	Login                   string      `json:"login"`
	Id                      int         `json:"id"`
	NodeId                  string      `json:"node_id"`
	AvatarUrl               string      `json:"avatar_url"`
	GravatarId              string      `json:"gravatar_id"`
	Url                     string      `json:"url"`
	HtmlUrl                 string      `json:"html_url"`
	FollowersUrl            string      `json:"followers_url"`
	FollowingUrl            string      `json:"following_url"`
	GistsUrl                string      `json:"gists_url"`
	StarredUrl              string      `json:"starred_url"`
	SubscriptionsUrl        string      `json:"subscriptions_url"`
	OrganizationsUrl        string      `json:"organizations_url"`
	ReposUrl                string      `json:"repos_url"`
	EventsUrl               string      `json:"events_url"`
	ReceivedEventsUrl       string      `json:"received_events_url"`
	Type                    string      `json:"type"`
	SiteAdmin               bool        `json:"site_admin"`
	Name                    string      `json:"name"`
	Company                 string      `json:"company"`
	Blog                    string      `json:"blog"`
	Location                string      `json:"location"`
	Email                   string      `json:"email"`
	Hireable                bool        `json:"hireable"`
	Bio                     string      `json:"bio"`
	TwitterUsername         interface{} `json:"twitter_username"`
	PublicRepos             int         `json:"public_repos"`
	PublicGists             int         `json:"public_gists"`
	Followers               int         `json:"followers"`
	Following               int         `json:"following"`
	CreatedAt               time.Time   `json:"created_at"`
	UpdatedAt               time.Time   `json:"updated_at"`
	PrivateGists            int         `json:"private_gists"`
	TotalPrivateRepos       int         `json:"total_private_repos"`
	OwnedPrivateRepos       int         `json:"owned_private_repos"`
	DiskUsage               int         `json:"disk_usage"`
	Collaborators           int         `json:"collaborators"`
	TwoFactorAuthentication bool        `json:"two_factor_authentication"`
	Plan                    struct {
		Name          string `json:"name"`
		Space         int    `json:"space"`
		Collaborators int    `json:"collaborators"`
		PrivateRepos  int    `json:"private_repos"`
	} `json:"plan"`
}

func (idp *GithubIdProvider) GetUserInfo(token *oauth2.Token) (*UserInfo, error) {
	req, err := http.NewRequest("GET", "https://api.github.com/user", nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Authorization", "token "+token.AccessToken)
	resp, err := idp.Client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var githubUserInfo GitHubUserInfo
	err = json.Unmarshal(body, &githubUserInfo)
	if err != nil {
		return nil, err
	}

	userInfo := UserInfo{
		Id:          strconv.Itoa(githubUserInfo.Id),
		Username:    githubUserInfo.Login,
		DisplayName: githubUserInfo.Name,
		Email:       githubUserInfo.Email,
		AvatarUrl:   githubUserInfo.AvatarUrl,
	}
	return &userInfo, nil
}

func (idp *GithubIdProvider) postWithBody(body interface{}, url string) ([]byte, error) {
	bs, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	r := strings.NewReader(string(bs))
	req, _ := http.NewRequest("POST", url, r)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	resp, err := idp.Client.Do(req)
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
