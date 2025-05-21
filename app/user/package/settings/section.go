package settings

type ServerConfig struct {
	Mode     string `mapstructure:"mode"`
	GinMode  string `mapstructure:"gin_mode"`
	GRPCPort int    `mapstructure:"grpc_port"`
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

type Locale struct {
	VIPath string `mapstructure:"locale_vi_path"`
	ENPath string `mapstructure:"locale_en_path"`
}

type ConsumeTopicsConfig struct {
	Key               string `mapstructure:"key"`
	Topic             string `mapstructure:"topic"`
	NumberConsumer    int    `mapstructure:"number_consumer"`
	NumberDLQConsumer int    `mapstructure:"number_dlq_consumer"`
	Handler           string `mapstructure:"handler"`
	DLQ               bool   `mapstructure:"dlq"`
}

type KafkaConsumerGroups struct {
	PrefixGroup   string                 `mapstructure:"prefix_group"`
	OffsetReset   string                 `mapstructure:"offset_reset"`
	ConsumeTopics []*ConsumeTopicsConfig `mapstructure:"consume_topics"`
}

type Kafka struct {
	Servers        string                 `mapstructure:"servers"`
	ConsumerGroups []*KafkaConsumerGroups `mapstructure:"consume_groups"`
}

type Config struct {
	Server         ServerConfig   `mapstructure:"server"`
	LogConfig      LogConfig      `mapstructure:"log"`
	SecurityConfig SecurityConfig `mapstructure:"security"`
	PostgresConfig PostgresConfig `mapstructure:"postgres"`
	Locale         Locale         `mapstructure:"locale"`
	Kafka          Kafka          `mapstructure:"kafka"`
}
