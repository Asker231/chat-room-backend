package db

import (
	"fmt"

	"github.com/Asker231/chat-room-backend.git/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DataBase struct{
	Db *gorm.DB
}

func InitDataBase(conf *config.DataBaseConfig)(DataBase){
	db,err := gorm.Open(postgres.Open(conf.DNS),&gorm.Config{})
	if err != nil{
		fmt.Println(err.Error())
	}
	return DataBase{
		Db: db,
	}
}