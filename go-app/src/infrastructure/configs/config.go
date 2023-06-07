package configs

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

var Config *conf

type Server struct {
	Common Common
	DB     DB
	Mail   Mail
}

type Common struct {
	JwtSecretKey   string `required:"true" envconfig:"JWT_SECRET_KEY"`
	SupportAddress string `required:"true" envconfig:"SUPPORT_ADDRESS"`
}

type DB struct {
	Host     string `required:"true" envconfig:"DB_HOST"`
	Port     string `required:"true" envconfig:"DB_PORT"`
	Name     string `required:"true" envconfig:"DB_NAME"`
	User     string `required:"true" envconfig:"DB_USER"`
	Password string `required:"true" envconfig:"DB_PASSWORD"`
}

type Mail struct {
	Host     string `required:"true" envconfig:"MAIL_HOST"`
	Port     string `required:"true" envconfig:"MAIL_PORT"`
	User     string `required:"true" envconfig:"MAIL_USER"`
	Password string `required:"true" envconfig:"MAIL_PASSWORD"`
}

type conf struct {
	Common Common
	DB     DB
	Mail   Mail
}

func InitConfig(fileName string) (*conf, error) {
	envPath := GetEnvFilePath("../../../" + fileName)

	if err := godotenv.Load(envPath); err != nil {
		fmt.Printf("%v", err.Error())
		os.Exit(1)
	}

	server := Server{}
	if err := envconfig.Process("server", &server); err != nil {
		return nil, err
	}

	Config = &conf{
		Common: server.Common,
		DB:     server.DB,
		Mail:   server.Mail,
	}
	return Config, nil
}

func GetEnvFilePath(relativePath string) string {
	_, currentFile, _, ok := runtime.Caller(0)
	if !ok {
		fmt.Fprintf(os.Stderr, "Unable to identify current directory")
		os.Exit(1)
	}
	basePath := filepath.Dir(currentFile)

	return filepath.Join(basePath, relativePath)
}
