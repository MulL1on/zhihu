package config

import "time"

type Middleware struct {
	Cors      *CORS      `mapstructure:"cors" yaml:"cors"`
	Jwt       *Jwt       `mapstructure:"jwt" yaml:"jwt"`
	RateLimit *RateLimit `mapstructure:"rateLimit" yaml:"rateLimit"`
}

type CORSWhitelist struct {
	AllowOrigin      string `mapstructure:"allowOrigin"      yaml:"allowOrigin"`
	AllowPath        string `mapstructure:"allowPath"        yaml:"AllowPath"`
	AllowMethods     string `mapstructure:"allowMethods"     yaml:"allowMethods"`
	AllowHeaders     string `mapstructure:"allowHeaders"     yaml:"allowHeaders"`
	ExposeHeaders    string `mapstructure:"exposeHeaders"    yaml:"exposeHeaders"`
	AllowCredentials bool   `mapstructure:"allowCredentials" yaml:"allowCredentials"`
}

type Jwt struct {
	SecretKey   string `mapstructure:"secretKey" yaml:"secretKey"`
	ExpiresTime int64  `mapstructure:"expiresTime" yaml:"expiresTime"`
	BufferTime  int64  `mapstructure:"bufferTime" yaml:"bufferTime"`
	Issuer      string `mapstructure:"issuer" yaml:"issuer"`
}

type CORS struct {
	Mode      string          `mapstructure:"mode" yaml:"mode"`
	Whitelist []CORSWhitelist `mapstructure:"whitelist" yaml:"whitelist"`
}

type RateLimit struct {
	Capacity     int64  `mapstructure:"capacity" yaml:"capacity"`
	Quantum      int64  `mapstructure:"quantum" yaml:"quantum"`
	FillInterval string `mapstructure:"fillInterval" yaml:"fillInterval"`
}

func (r *RateLimit) GetFillInterval(fillInterval string) time.Duration {
	t, _ := time.ParseDuration(fillInterval)
	return t
}
