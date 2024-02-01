package dto

type TokenResponse struct {
	Code         int    `json:"-"`
	Msg          string `json:"msg"`
	TokenType    string `json:"token_type"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpireIn     int    `json:"expire_in"`
	ExpireAt     int    `json:"expire_at"`
}
