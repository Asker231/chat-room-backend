package main

import (
	"os"

	"github.com/Asker231/chat-room-backend.git/internal/user"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func main(){
	if err := godotenv.Load(".env");err != nil{
		return
	}

	db,err := gorm.Open(postgres.Open(os.Getenv("DNS")),&gorm.Config{})
	if err != nil{
		return
	}
	db.AutoMigrate(user.User{})

}