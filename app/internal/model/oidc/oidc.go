package oidc

type TokenResponse struct {
	TokenType string `json:"token_type"`
	ExpiresIn int64  `json:"expires_in"`
	IdToken   string `json:"id_token"`
}
