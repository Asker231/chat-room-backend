package req

import (
	"encoding/json"
	"net/http"

	"github.com/Asker231/chat-room-backend.git/pkg/res"
	"github.com/go-playground/validator/v10"
)


func HandleBody[T any](w http.ResponseWriter,r *http.Request)(*T,error){
	var payload T
	if err := json.NewDecoder(r.Body).Decode(&payload);err != nil{
		res.ResponseBody(w,err.Error(),http.StatusBadRequest)
		return nil,err
	}
	validate := validator.New()
	if err := validate.Struct(payload);err != nil{
		res.ResponseBody(w,err.Error(),401)
		return nil,err
	}
	return &payload,nil
}