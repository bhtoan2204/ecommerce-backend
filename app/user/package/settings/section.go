package settings

type ServerConfig struct {
	Mode    string `mapstructure:"mode"`
	GinMode string `mapstructure:"gin_mode"`
}

type SecurityConfig struct {
	JWTAccessSecret      string `mapstructure:"jwt_access_secret"`
	JWTRefreshSecret     string `mapstructure:"jwt_refresh_secret"`
	JWTAccessExpiration  int    `mapstructure:"jwt_access_expiration"`
	JWTRefreshExpiration int    `mapstructure:"jwt_refresh_expiration"`
}

type LogConfig struct {
	LogLevel   string `mapstructure:"log_level"`
	FilePath   string `mapstructure:"file_path"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxBackups int    `mapstructure:"max_backups"`
	MaxAge     int    `mapstructure:"max_age"`
	Compress   bool   `mapstructure:"compress"`
}

type Config struct {
	Server         ServerConfig   `mapstructure:"server"`
	LogConfig      LogConfig      `mapstructure:"log"`
	SecurityConfig SecurityConfig `mapstructure:"security"`
}
