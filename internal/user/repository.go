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
func(repo *UserRepo)AddUser(user *User)(*User,error){
	result := repo.DataBase.Create(user)

	if result.Error != nil{
		return nil,result.Error
	}
	return user,nil
}

//fyndByEmailUser db
func(repo *UserRepo)FindUserByEmail(email string)(*User,error){
	var user User
	result := repo.DataBase.First(&user,"email = ?",email)
	if result.Error != nil{
		return nil,result.Error
	}
	return &user,nil
}