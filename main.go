package main

import (
	"golang_restful_api/app"
	"golang_restful_api/controller"
	"golang_restful_api/repository"
	"golang_restful_api/service"

	"github.com/go-playground/validator/v10"
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
	router.GET("/api/categories/:categoryId", categoryController.FindByid)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)
}