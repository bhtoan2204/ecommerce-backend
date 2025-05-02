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

type PostgresConfig struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	Username     string `mapstructure:"username"`
	Password     string `mapstructure:"password"`
	Database     string `mapstructure:"database"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
	MaxLifetime  int    `mapstructure:"max_lifetime"`
}

type Config struct {
	Server         ServerConfig   `mapstructure:"server"`
	LogConfig      LogConfig      `mapstructure:"log"`
	SecurityConfig SecurityConfig `mapstructure:"security"`
	PostgresConfig PostgresConfig `mapstructure:"postgres"`
}
