package setting

type Config struct {
	Server ServerSetting `mapstructure:"server"`
	Logger LoggerSetting `mapstructure:"logger"`
	PostgreSQL DatabaseSetting `mapstructure:"postgresql"`
	Redis RedisSetting `mapstructure:"redis"`
}

type ServerSetting struct {
	Port string `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
	SecretKey string `mapstructure:"secret_key"`
	RateLimit int `mapstructure:"rate_limit"`
	RateLimitDuration int `mapstructure:"rate_limit_duration"`
}

type LoggerSetting struct {
	LogLevel string `mapstructure:"log_level"`
	FileName string `mapstructure:"file_log_name"`
	MaxSize int `mapstructure:"max_size"`
	MaxBackups int `mapstructure:"max_backups"`
	MaxAge int `mapstructure:"max_age"`
	Compress bool `mapstructure:"compress"`
}

type DatabaseSetting struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
	User string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DbName string `mapstructure:"dbname"`
}

type RedisSetting struct {
	User string `mapstructure:"user"`
	Addr string `mapstructure:"addr"`
	Password string `mapstructure:"password"`
	DB int `mapstructure:"db"`
}