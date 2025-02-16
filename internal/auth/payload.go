package auth

type(
	RegisterRequest struct{
		Name string `json:"name" validate:"required"`
		Email string `josn:"email" validate:"email,required"`
		Password string `json:"password" validate:"required"`
	} 
	RegisterResponse struct{
		Token string `json:"token"`
	}


	LoginRequest struct{
		Email string `josn:"email" validate:"email,required"`
		Password string `json:"password" validate:"required"`
	}
	LoginResponse struct{
		Token string `json:"token"`
	}
)