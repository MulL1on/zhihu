package config

type Config struct {
	Logger     *Logger     `mapstructure:"logger" yaml:"logger"`
	Database   *Database   `mapstructure:"database" yaml:"database"`
	Middleware *Middleware `mapstructure:"middleware" yaml:"middleware"`
	App        *App        `mapstructure:"app" yaml:"app"`
	Server     *Server     `mapstructure:"server" yaml:"server"`
}
