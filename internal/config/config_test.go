package config

import (
	"os"
	"reflect"
	"testing"
	"time"
)

func TestInitConfig(t *testing.T) {
	type env struct {
		mongodbURL           string
		redisAddress         string
		emailServiceName     string
		emailServiceAddress  string
		emailServicePassword string
		secretKey            string
		codedSalt            string
		appEnv               string
		httpHost             string
		folderAccounts       string
		folderPythonScripts  string
	}

	type args struct {
		path string
		env  env
	}

	setEnv := func(env env) {
		os.Setenv("MONGODB_URL", env.mongodbURL)
		os.Setenv("REDIS_ADDRESS", env.redisAddress)
		os.Setenv("EMAIL_SERVICE_NAME", env.emailServiceName)
		os.Setenv("EMAIL_SERVICE_ADDRESS", env.emailServiceAddress)
		os.Setenv("EMAIL_SERVICE_PASSWORD", env.emailServicePassword)
		os.Setenv("SECRET_KEY", env.secretKey)
		os.Setenv("CODE_SALT", env.codedSalt)
		os.Setenv("ENV", env.appEnv)
		os.Setenv("HTTP_HOST", env.httpHost)
		os.Setenv("FOLDER_ACCOUNTS", env.folderAccounts)
		os.Setenv("FOLDER_PYTHON_SCRIPTS", env.folderPythonScripts)
	}

	tests := []struct {
		name    string
		args    args
		want    *Config
		wantErr bool
	}{
		{
			name: "test config",
			args: args{
				path: "fixtures",
				env: env{
					mongodbURL:           "mongodb://127.0.0.1:27017",
					redisAddress:         "0.0.0.0:6379",
					emailServiceName:     "Service",
					emailServiceAddress:  "service@gmail.com",
					emailServicePassword: "qwerty123",
					secretKey:            "secret_key",
					codedSalt:            "code_salt",
					appEnv:               "local",
					httpHost:             "localhost",
					folderAccounts:       "accounts",
					folderPythonScripts:  "python",
				},
			},
			want: &Config{
				Environment: "local",
				Mongo: MongoConfig{
					URL:          "mongodb://127.0.0.1:27017",
					Username:     "vanya",
					DatabaseName: "services",
				},
				Redis: RedisConfig{
					Address: "0.0.0.0:6379",
				},
				Email: EmailConfig{
					ServiceName:     "Service",
					ServiceAddress:  "service@gmail.com",
					ServicePassword: "qwerty123",
					Templates: EmailTemplates{
						LoginNotification: "./templates/login_notification.html",
					},
					Subjects: EmailSubjects{
						LoginNotification: "Уведомление о входе в аккаунт",
					},
				},
				Auth: AuthConfig{
					JWT: JWTConfig{
						AccessTokenTTL:  time.Minute * 15,
						RefreshTokenTTL: time.Hour * 720,
					},
					SecretCodeLifetime:     time.Minute * 5,
					VerificationCodeLength: 6,
					SecretKey:              "secret_key",
					CodeSalt:               "code_salt",
				},
				HTTP: HTTPConfig{
					Host:               "localhost",
					Port:               "8080",
					MaxHeaderMegabytes: 1,
					ReadTimeout:        time.Second * 10,
					WriteTimeout:       time.Second * 10,
				},
				SMTP: SMTPConfig{
					Host: "smtp.gmail.com",
					Port: 587,
				},
				Folders: FoldersConfig{
					Accounts:      "accounts",
					PythonScripts: "python",
				},
			},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			setEnv(testCase.args.env)

			got, err := InitConfig(testCase.args.path)
			if (err != nil) != testCase.wantErr {
				t.Errorf("InitConfig() error = %v, wantErr %v", err, testCase.wantErr)

				return
			}
			if !reflect.DeepEqual(got, testCase.want) {
				t.Errorf("InitConfig() got = %v, want = %v", got, testCase.want)
			}
		})
	}
}
