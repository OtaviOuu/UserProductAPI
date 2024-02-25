package configs

import (
	"github.com/go-chi/jwtauth"
)

type config struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBHost        string `mapstructure:"DB_HOST"`
	DBPort        int    `mapstructure:"DB_PORT"`
	DBuser        string `mapstructure:"DB_USER"`
	DBPassword    string `mapstructure:"DB_PASSWORD"`
	DBName        string `mapstructure:"DB_NAME"`
	WebServerPort string `mapstructure:"WEB_SERVER_PORT"`
	JWTSecret     string `mapstructure:"JWT_SECRETS"`
	JWTExperesIn  int    `mapstructure:"JWT_EXPIRESIN"`
	TokenAuth     *jwtauth.JWTAuth
}

func LoadConfig() *config {

	cfg := &config{
		DBDriver:      "mysql",
		DBHost:        "localhost",
		DBPort:        3306,
		DBuser:        "root",
		DBPassword:    "root",
		DBName:        "dbApi",
		WebServerPort: ":8000",
		JWTSecret:     "secret",
		JWTExperesIn:  300,
	}

	cfg.TokenAuth = jwtauth.New("HS256", []byte("secret"), nil)

	return cfg
}
