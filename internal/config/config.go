package config

import (
	"github.com/caarlos0/env/v10"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

const (
	EnvLocal = "local"
	EnvProd  = "prod"
)

type (
	Config struct {
		Environment        string         `env:"ENV"`
		Postgres           PostgresConfig `mapstructure:"postgresql"`
		Mongo              MongoConfig    `mapstructure:"mongodb"`
		Redis              RedisConfig
		HTTP               HTTPConfig  `mapstructure:"http"`
		Auth               AuthConfig  `mapstructure:"auth"`
		SMTP               SMTPConfig  `mapstructure:"smtp"`
		Email              EmailConfig `mapstructure:"email"`
		UrlListenOAuthCode string      `env:"URL_LISTEN_OAUTH_CODE"`
		Folders            FoldersConfig
	}

	PostgresConfig struct {
		URL          string        `env:"POSTGRESQL_URL"`
		MigrationURL string        `env:"MIGRATION_URL"`
		MaxAttempts  int           `mapstructure:"max_attempts"`
		MaxDelay     time.Duration `mapstructure:"max_delay"`
	}

	MongoConfig struct {
		URL          string `env:"MONGODB_URL"`
		Username     string `mapstructure:"username"`
		DatabaseName string `mapstructure:"databaseName"`
	}

	RedisConfig struct {
		Address string `env:"REDIS_ADDRESS"`
	}

	EmailConfig struct {
		ServiceName     string         `env:"EMAIL_SERVICE_NAME"`
		ServiceAddress  string         `env:"EMAIL_SERVICE_ADDRESS"`
		ServicePassword string         `env:"EMAIL_SERVICE_PASSWORD"`
		Templates       EmailTemplates `mapstructure:"templates"`
		Subjects        EmailSubjects  `mapstructure:"subjects"`
	}

	EmailTemplates struct {
		VerifyEmail       string `mapstructure:"verify_email"`
		LoginNotification string `mapstructure:"login_notification"`
	}

	EmailSubjects struct {
		VerifyEmail       string `mapstructure:"verify_email"`
		LoginNotification string `mapstructure:"login_notification"`
	}

	AuthConfig struct {
		JWT                    JWTConfig     `mapstructure:"jwt"`
		SercetCodeLifetime     time.Duration `mapstructure:"sercetCodeLifetime"`
		VerificationCodeLength int           `mapstructure:"verificationCodeLength"`
		SecretKey              string        `env:"SECRET_KEY"`
		CodeSalt               string        `env:"CODE_SALT"`
	}

	JWTConfig struct {
		AccessTokenTTL  time.Duration `mapstructure:"accessTokenTTL"`
		RefreshTokenTTL time.Duration `mapstructure:"refreshTokenTTL"`
	}

	HTTPConfig struct {
		Host               string        `env:"HTTP_HOST"`
		Port               string        `mapstructure:"port"`
		MaxHeaderMegabytes int           `mapstructure:"maxHeaderBytes"`
		ReadTimeout        time.Duration `mapstructure:"readTimeout"`
		WriteTimeout       time.Duration `mapstructure:"writeTimeout"`
	}

	SMTPConfig struct {
		Host string `mapstructure:"host"`
		Port int    `mapstructure:"port"`
	}

	FoldersConfig struct {
		Accounts      string `env:"FOLDER_ACCOUNTS"`
		Channels      string `env:"FOLDER_CHANNELS"`
		PythonScripts string `env:"FOLDER_PYTHON_SCRIPTS"`
	}
)

func InitConfig(configPath string) (*Config, error) {
	if err := parseConfigFile(configPath); err != nil {
		return nil, err
	}

	var cfg Config

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	if os.Getenv("APP_ENV") == "local" {
		if err := godotenv.Load(); err != nil {
			return nil, err
		}
	}

	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func parseConfigFile(folder string) error {
	viper.AddConfigPath(folder)
	viper.SetConfigName("main")

	return viper.ReadInConfig()
}
