package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Kontrak CategoryController mendefinisikan metode yang harus diimplementasikan oleh controller kategori
type CategoryController interface {
	// Metode-metode dibawah ini akan menangani permintaan HTTP untuk operasi CRUD pada kategori
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindByid(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}