package helper

import (
	"encoding/json"
	"net/http"
)

func ReadFromRequestBody(request *http.Request, result interface{}) {
	decoder := json.NewDecoder(request.Body) // Membuat decoder untuk membaca JSON dari body request
	err := decoder.Decode(result)            // Mengikatkan data dari request body ke struct result
	PanicIfError(err)                        // Menangani error jika terjadi saat decoding
}

func WriteToResponse(writer http.ResponseWriter, response interface{}) {
	writer.Header().Add("Content-Type", "application/json") 	// Menetapkan header Content-Type sebagai JSON
	encoder := json.NewEncoder(writer)                       	// Membuat encoder untuk menulis JSON ke response writer
	err := encoder.Encode(response)                          	// Mengencode data ke dalam format JSON dan menulisnya ke response writer
	PanicIfError(err)                                        	// Menangani error jika terjadi saat encoding
}