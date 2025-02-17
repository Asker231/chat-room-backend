package main

import (
	"log"
	"net/http"

	"github.com/Asker231/chat-room-backend.git/config"
	"github.com/Asker231/chat-room-backend.git/internal/auth"
	"github.com/Asker231/chat-room-backend.git/internal/user"
	"github.com/Asker231/chat-room-backend.git/pkg/db"

	"github.com/rs/cors"
)


func main() {

	//init config
	cnf:= config.InitConfig() 
	//init db
	dataBase := db.InitDataBase(&cnf.Db)
	//repo
	userRepo := user.NewRepoUser(dataBase.Db)

	//service
	service:= auth.NewAuthService(userRepo)

	//init Handler
	router := http.NewServeMux()
	//handler
	auth.NewAuthHandler(router,cnf,service)
	//cors
	handler := cors.AllowAll().Handler(router)

	//init server 
	server := http.Server{
		Addr: ":8080",
		Handler: handler,
	}

	log.Panic(server.ListenAndServe())
}