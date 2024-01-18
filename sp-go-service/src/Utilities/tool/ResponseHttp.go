package tool

import (
	"encoding/json"
	"net/http"
	"sp-go-service/src/Resources/global"
)

func ResponseHttp(w http.ResponseWriter) {
	GetBodyWithLastCode()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(global.ResponseBody)
}
