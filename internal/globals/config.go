package globals

import (
	"log"
	"os"

	"github.com/goccy/go-yaml"
	"github.com/todalist/app/internal/common"
	"go.uber.org/zap"
)

var (
	CONF *AppConfig
)

type AppConfig struct {
	Server *ServerConfig
	DB     *DBConfig
	Auth   *AuthenticationConfig
	Redis  *RedisConfig
}

type RedisConfig struct {
	Host     string
	Port     uint16
	Db       int
	Password string
}

type ServerConfig struct {
	// Name       string
	Port       uint16
	PathPrefix string `yaml:"pathPrefix"`
	Cors       CorsConfig
}

type DBConfig struct {
	Host     string
	User     string
	Password string
	Database string
	Port     uint16
}

type AuthenticationConfig struct {
	Jwt       *JwtConfig
	WhiteList []string `yaml:"whiteList"`
}

type JwtConfig struct {
	JwtExpireSec int64  `yaml:"jwtExpireSec"`
	JwtSecret    string `yaml:"jwtSecret"`
	JwtIssuer    string `yaml:"jwtIssuer"`
}

type CorsConfig struct {
	Enable  bool
	Origins []string
}

func MustLoad() *AppConfig {
	file_path := os.Getenv(common.APP_CONFIG_ENV_KEY)
	if file_path == "" {
		log.Fatalf("config file is require. specific via '%s' env", common.APP_CONFIG_ENV_KEY)
	}
	config_content, err := os.ReadFile(file_path)
	if err != nil {
		log.Fatalf("read config error: %v", err)
	}
	config := new(AppConfig)
	if err = yaml.Unmarshal(config_content, config); err != nil {
		log.Fatalf("unmarshal config error: %v", err)
	}
	// set global configuration object
	CONF = config
	LOG.Debug("configuration loaded",
		zap.String("file_path", file_path),
		zap.Any("config", CONF))
	return config
}

func IsDev() bool {
	return os.Getenv(common.APP_DEVELOPMENT_ENV_KEY) == "1"
}

func IsProd() bool {
	return !IsDev()
}
