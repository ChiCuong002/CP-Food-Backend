package setting

type Config struct {
	Server ServerSetting `mapstructure:"server"`
	Logger LoggerSetting `mapstructure:"logger"`
	PostgreSQL DatabaseSetting `mapstructure:"postgresql"`
}

type ServerSetting struct {
	Port string `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
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