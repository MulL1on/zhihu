package config

type Upload struct {
	AccessKey string `mapstructure:"accessKey" yaml:"accessKey"`
	SecretKey string `mapstructure:"secretKey" yaml:"secretKey"`
	Bucket    string `mapstructure:"bucket" yaml:"bucket"`
	Server    string `mapstructure:"server" yaml:"server"`
}