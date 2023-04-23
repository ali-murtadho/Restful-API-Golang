package utils

import (
	"encoding/json"
	"net/http"
)

func ResponseJSON(rw http.ResponseWriter, p interface{}, status int) {
	ChangeToByte, err := json.Marshal(p)
	rw.Header().Set("Content-Type","application/json")

	if err != nil {
		http.Error(rw, "Error", http.StatusBadRequest)	
	}
	rw.Header().Set("Content-Type","application/json")
	rw.WriteHeader(status)
	rw.Write([]byte(ChangeToByte))


}