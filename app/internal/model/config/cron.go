package config

type Cron struct {
	ScanCounterSpec   string `mapstructure:"scanCounterSpec" yaml:"scanCounterSpec"`
	ScanCheckDiggSpec string `mapstructure:"scanCheckDiggSpec" yaml:"scanCheckDiggSpec"`
}
