package config

type Config struct {
	Logger     *Logger     `mapstructure:"logger" yaml:"logger"`
	Database   *Database   `mapstructure:"database" yaml:"database"`
	Middleware *Middleware `mapstructure:"middleware" yaml:"middleware"`
	App        *App        `mapstructure:"app" yaml:"app"`
	Server     *Server     `mapstructure:"server" yaml:"server"`
	Snowflake  *Snowflake  `mapstructure:"snowflake" yaml:"snowflake"`
	Cron       *Cron       `mapstructure:"cron" yaml:"cron"`
	Upload     *Upload     `mapstructure:"upload" yaml:"upload"`
}
