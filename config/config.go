package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Database database
	Redis    redis
}

type redis struct {
	Host string `envconfig:"APP_REDIS_HOST" default:"localhost"`
	Port string `envconfig:"APP_REDIS_PORT" default:"6379"`
}

type database struct {
	Driver    string `envconfig:"APP_DB_DRIVER" default:"postgres"`
	Host      string `envconfig:"APP_DB_HOST" default:"localhost"`
	Port      string `envconfig:"APP_DB_PORT" default:"5432"`
	Name      string `envconfig:"APP_DB_NAME" default:"opt-auth"`
	Username  string `envconfig:"APP_DB_USERNAME" default:"opt-auth"`
	Password  string `envconfig:"APP_DB_PASSWORD" default:"opt-auth"`
	EnableSSL bool   `envconfig:"APP_DB_ENABLE_SSL" default:"false"`
}

func (db database) ConnStr() string {
	sslMode := "disable"
	if db.EnableSSL {
		sslMode = "enable"
	}

	switch db.Driver {
	case "postgres":
		return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", db.Host, db.Port, db.Username, db.Password, db.Name, sslMode)

	case "sqlite3":
		return "file:opt-auth.db?cache=shared&mode=rwc"

	default:
		return ""
	}
}

func (r redis) Addr() string {
	return r.Host + ":" + r.Port
}

func Environ() (Config, error) {
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return Config{}, err
	}
	return cfg, nil
}
