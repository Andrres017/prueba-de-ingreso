package common

import (
	"net/http"
)

func SendResponse(write http.ResponseWriter, status int, data []byte) {

	write.Header().Set("Content-Type", "application/json")
	write.WriteHeader(status)
	write.Write(data)
}

func SendError(write http.ResponseWriter, status int) {
	data := []byte(`{}`)
	write.Header().Set("Content-Type", "application/json")
	write.WriteHeader(status)
	write.Write(data)
}
