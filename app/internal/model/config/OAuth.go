package config

type GithubOAuth struct {
	ClientId     string `mapstructure:"clientId" yaml:"clientId" json:"client_id"`
	ClientSecret string `mapstructure:"clientServer" yaml:"clientSecret" json:"client_secret"`
	RedirectUri  string `mapstructure:"redirectUri" yaml:"redirectUri" json:"redirect_uri"`
}
