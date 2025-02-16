package main

import (
	"log"
	"net/http"

	"github.com/Asker231/chat-room-backend.git/config"
	"github.com/Asker231/chat-room-backend.git/pkg/db"
)


func main() {

	//init config
	conf := config.InitConfig() 
	//init db
	_ = db.InitDataBase(&conf.Db)
	//init Handler
	router := http.NewServeMux()

	//init server 
	server := http.Server{
		Addr: ":80",
		Handler: router,
	}

	log.Panic(server.ListenAndServe())
}