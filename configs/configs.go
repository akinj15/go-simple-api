package configs

import (
	"github.com/go-chi/jwtauth/v5"
	"github.com/spf13/viper"
)

var cfg *Config
var tokenAuth *jwtauth.JWTAuth

type Config struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBHOST        string `mapstructure:"DB_HOST"`
	DBPort        string `mapstructure:"DB_PORT"`
	DBUser        string `mapstructure:"DB_USER"`
	DBPassword    string `mapstructure:"DB_PASSWORD"`
	DBName        string `mapstructure:"DB_NAME"`
	WebServerPort string `mapstructure:"WEB_SERVER_PORT"`
	JWTSecret     string `mapstructure:"JWT_SECRET"`
	JWTExpireIn   int    `mapstructure:"JWT_EXPIRE_IN"`
	TokenAuth     *jwtauth.JWTAuth
}

func LoadConfig(path string) (*Config, error) {
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}
	tokenAuth = jwtauth.New("HS256", []byte(cfg.JWTSecret), nil) // replace with secret key

	cfg.TokenAuth = tokenAuth

	return cfg, nil
}
