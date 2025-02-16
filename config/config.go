package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)


type(
	AppConfig struct{
		Auth AuthConfig
		Db DataBaseConfig
	}
	AuthConfig struct{
		SECRET string
	}
	DataBaseConfig struct{
		DNS string
	}
)

func InitConfig()*AppConfig{
	if err := godotenv.Load(".env");err != nil{
		fmt.Println(err.Error())
		return nil
	}
	return &AppConfig{
		Auth: AuthConfig{
			SECRET: os.Getenv("SECRET"),
		},
		Db :DataBaseConfig{
			DNS:os.Getenv("DNS"),
		},
	}
}