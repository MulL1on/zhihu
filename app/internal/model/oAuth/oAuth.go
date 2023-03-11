package oAuth

type GithubOAuthAc struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}

type GithubUserInfo struct {
	Login    string `json:"login"`
	GithubId int64  `json:"id"`
}
