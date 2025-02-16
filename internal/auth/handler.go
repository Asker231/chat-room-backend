package auth

import (
	"fmt"
	"net/http"

	"github.com/Asker231/chat-room-backend.git/config"
	"github.com/Asker231/chat-room-backend.git/pkg/req"
	"github.com/Asker231/chat-room-backend.git/pkg/res"
)

type AuthHandler struct{
	Config *config.AppConfig
}
func NewAuthHandler(router *http.ServeMux,conf *config.AppConfig){
	handler := &AuthHandler{
		Config: conf,
	}

	router.HandleFunc("POST /auth/register",handler.Register())
	router.HandleFunc("POST /auth/login",handler.Login())
}


func(handler *AuthHandler)Register()http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		body,err := req.HandleBody[RegisterRequest](w,r)
		if err != nil{
			res.ResponseBody(w,err.Error(),404)
			return
		}	
		res.ResponseBody(w,RegisterResponse{
			Token: "199823231998",
		},200)

		fmt.Println(body)
	}
}
func(handler *AuthHandler)Login()http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		body,err := req.HandleBody[LoginRequest](w,r)
		if err != nil{
			res.ResponseBody(w,err.Error(),404)
			return
		}	
		res.ResponseBody(w,LoginResponse{
			Token: "199823231998",
		},200)

		fmt.Println(body)
	}
}




