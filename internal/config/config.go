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
database:
  name: conduit
  host: localhost
  port: 5432
  username: tinhtt
  password: tinhtt
http_server:
  port: :8080
  jwt_secret: verysecret
  jwt_duration: 168h
`)

type Config struct {
	Foo      string `mapstructure:"foo"`
	Mode     Mode   `mapstructure:"mode"`
	Database struct {
		Name     string `mapstructure:"name"`
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
	} `mapstructure:"database"`

	HTTPServer struct {
		Port        string        `mapstructure:"port"`
		JWTSecret   string        `mapstructure:"jwt_secret"`
		JWTDuration time.Duration `mapstructure:"jwt_duration"`
	} `mapstructure:"http_server"`
}

func Load() (Config, error) {
	v := viper.New()
	v.SetConfigType("yaml")
	if err := v.ReadConfig(bytes.NewBuffer(rawConfig)); err != nil {
		return Config{}, err
	}

	v.AutomaticEnv()
	v.SetEnvPrefix("CONDUIT")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	var c Config
	if err := v.Unmarshal(&c); err != nil {
		return Config{}, err
	}

	return c, nil
}
