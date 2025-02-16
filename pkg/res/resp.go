package res

import (
	"encoding/json"
	"fmt"
	"net/http"
)


func ResponseBody(w http.ResponseWriter,data any,StatusCode int){
	err := json.NewEncoder(w).Encode(data)
	if  err != nil 	{
		fmt.Println(err.Error())
	}
	w.WriteHeader(StatusCode)
}