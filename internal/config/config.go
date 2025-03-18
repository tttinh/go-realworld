package config

import (
	"bytes"
	"strings"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/spf13/viper"
)

type Mode string

const (
	Test    = "test"
	Debug   = "debug"
	Release = "release"
)

var rawConfig = []byte(`
mode: debug

migration:
  auto: true
  fresh: false
  path: db/postgres/migration

database:
  name: conduit
  host: localhost
  port: 5432
  username: tinhtt
  password: tinhtt

http:
  port: :8080
  jwt_secret: verysecret
  jwt_duration: 168h
`)

type Config struct {
	Mode Mode `mapstructure:"mode" json:"mode"`

	Migration struct {
		Path  string `mapstructure:"path" json:"path"`
		Auto  bool   `mapstructure:"auto" json:"auto"`
		Fresh bool   `mapstructure:"fresh" json:"fresh"`
	} `mapstructure:"migration" json:"migration"`

	Database struct {
		Name     string `mapstructure:"name" json:"name"`
		Host     string `mapstructure:"host" json:"host"`
		Port     int    `mapstructure:"port" json:"port"`
		Username string `mapstructure:"username" json:"username"`
		Password string `mapstructure:"password" json:"password"`
	} `mapstructure:"database" json:"database"`

	HTTP struct {
		Port        string        `mapstructure:"port" json:"port"`
		JWTSecret   string        `mapstructure:"jwt_secret" json:"jwt_secret"`
		JWTDuration time.Duration `mapstructure:"jwt_duration" json:"jwt_duration"`
	} `mapstructure:"http" json:"http"`
}

func Load() (*Config, error) {
	v := viper.New()
	v.SetConfigType("yaml")
	if err := v.ReadConfig(bytes.NewBuffer(rawConfig)); err != nil {
		return nil, err
	}

	v.AutomaticEnv()
	v.SetEnvPrefix("CONDUIT")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	var c Config
	if err := v.Unmarshal(&c); err != nil {
		return nil, err
	}

	return &c, nil
}
