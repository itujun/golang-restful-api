package web

type WebResponse struct {
	Code   int         `json:"code"`   // Kode status HTTP
	Status string      `json:"status"` // Deskripsi status HTTP
	Data   interface{} `json:"data"`   // Data yang ingin dikirimkan dalam respons
}