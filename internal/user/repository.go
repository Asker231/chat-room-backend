package user

import "gorm.io/gorm"

type UserRepo struct{
	DataBase *gorm.DB
}

func NewRepoUser(db *gorm.DB)*UserRepo{
	return &UserRepo{
		DataBase: db,
	}
}
//register user db

//fyndByEmailUser db
