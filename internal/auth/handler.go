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
	Service *AuthService
}
func NewAuthHandler(router *http.ServeMux,conf *config.AppConfig,service *AuthService){
	handler := &AuthHandler{
		Config: conf,
		Service: service,
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

		user,err := handler.Service.Register(body.Name,body.Email,body.Password)
		if err != nil{
			return
		}

		res.ResponseBody(w,user,201)

		// res.ResponseBody(w,RegisterResponse{
		// 	Token: "199823231998",
		// },200)
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




