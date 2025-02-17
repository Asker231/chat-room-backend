package auth

import (
	"github.com/Asker231/chat-room-backend.git/internal/user"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct{
	UserRepo *user.UserRepo
}


func NewAuthService(userRepo *user.UserRepo)*AuthService{
	return &AuthService{
		UserRepo: userRepo,
	}
}

func(serv *AuthService)Register(name,email,password string)(*user.User,error){
	pass , err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost);
	if err != nil{
		return nil,err
	}
	user := &user.User{
		Name: name,
		Email: email,
		Password: string(pass),
	}
	user ,err = serv.UserRepo.AddUser(user)
	if err != nil{
		return nil,err
	}
	return user,nil
}