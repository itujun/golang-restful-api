package main

import (
	"golang-restful-api/app"
	"golang-restful-api/controller"
	"golang-restful-api/exception"
	"golang-restful-api/helper"
	"golang-restful-api/middleware"
	"golang-restful-api/repository"
	"golang-restful-api/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql" // Import driver MySQL untuk database
	"github.com/julienschmidt/httprouter"
)

func main() {
	//  koneksikan ke database
	db := app.NewDB()
	// defer db.Close() // Pastikan koneksi database ditutup setelah digunakan
	
	//  initialisasi validator
	validate := validator.New()

	//  inisialisasi repository
	categoryRepository := repository.NewCategoryRepository()

	//  inisialisasi service
	categoryService := service.NewCategoryService(categoryRepository, db, validate)

	//  inisialisasi controller
	categoryController := controller.NewCategoryController(categoryService)

	//  inisialisasi router
	router := httprouter.New()

	//  menambahkan route untuk kategori
	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	router.PanicHandler = exception.ErrorHandler

	// Membuat Server
	server := http.Server{
		Addr:    "localhost:3000", 						// Alamat server
		Handler: middleware.NewAuthMiddleware(router), 	// Router yang menangani permintaan HTTP
	}

	err := server.ListenAndServe() 	// Memulai server dan mendengarkan permintaan
	helper.PanicIfError(err) 		// Menangani error jika terjadi saat memulai server
}