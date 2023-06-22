package configs

import (
	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

var cfg *conf

type conf struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBHost        string `mapstructure:"DB_HOST"`
	DBPort        string `mapstructure:"DB_PORT"`
	DBUser        string `mapstructure:"DB_USER"`
	DBPassword    string `mapstructure:"DB_PASSWORD"`
	DBName        string `mapstructure:"DB_NAME"`
	WebServerPort string `mapstructure:"DB_WEB_SERVER_PORT"`
	JWTSecret     string `mapstructure:"DB_JWT_SECRET"`
	JwtExpiresIn  int    `mapstructure:"DB_EXPIRES_IN"`
	TokenAuthKey  *jwtauth.JWTAuth
}

func LoadConfig(path string) *conf {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(path)
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = v.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}
	cfg.TokenAuthKey = jwtauth.New("HS256", []byte(cfg.JWTSecret), nil)
	return cfg
}
