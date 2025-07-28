package controller

import (
	"golang_restful_api/helper"
	"golang_restful_api/model/web"
	"golang_restful_api/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService // Menyatakan bahwa CategoryControllerImpl adalah implementasi dari CategoryService
}

// Fungsi untuk membuat instance baru dari CategoryControllerImpl
func NewCategoryController(categoryService service.CategoryService) CategoryController { 
	return &CategoryControllerImpl{
		CategoryService: categoryService, // Inisialisasi CategoryService
	}
}


func (controller *CategoryControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// Implementasi untuk membuat kategori

	// // 1. Membaca request body dan mengikatnya ke CategoryCreateRequest
	// decoder := json.NewDecoder(request.Body)
	// // mengikatkan data dari request body ke struct CategoryCreateRequest
	// categoryCreateRequest := web.CategoryCreateRequest{}
	// err := decoder.Decode(&categoryCreateRequest)
	// helper.PanicIfError(err)

	// 1. Dirapikan/Destrucring
	categoryCreateRequest := web.CategoryCreateRequest{}
	helper.ReadFromRequestBody(request, &categoryCreateRequest) // Membaca request body dan mengikatnya ke CategoryCreateRequest

	// 2. Melakukan validasi terhadap request

	// 3. Memanggil service untuk menyimpan kategori baru
	categoryResponse := controller.CategoryService.Create(request.Context(), categoryCreateRequest)
	webResponse := web.WebResponse{
		Code:   http.StatusCreated, 	// Kode status HTTP untuk Created
		Status: "OK",	 				// Deskripsi status HTTP
		Data:   categoryResponse,	 	// Data yang ingin dikirimkan dalam respons
	}

	// // 4. Mengembalikan response yang sesuai
	// writer.Header().Add("Content-Type", "application/json")	// Menetapkan header Content-Type sebagai JSON
	// encoder := json.NewEncoder(writer)						// Membuat encoder untuk menulis JSON ke response writer
	// err = encoder.Encode(webResponse)						// Mengencode webResponse ke dalam format JSON dan menulisnya ke response writer
	// helper.PanicIfError(err)								// Menangani error jika terjadi saat encoding

	// 4. Dirapikan/Destrucring
	helper.WriteToResponse(writer, webResponse) // Mengembalikan response yang sesuai dengan format JSON
}

func (controller *CategoryControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// Implementasi untuk memperbarui kategori

	// 1. Membaca request body dan mengikatnya ke CategoryCreateRequest
	categoryUpdateRequest := web.CategoryUpdateRequest{}
	helper.ReadFromRequestBody(request, &categoryUpdateRequest) // Membaca request body dan mengikatnya ke CategoryUpdateRequest

	categoryId := params.ByName("id") 	// Mengambil ID kategori dari parameter URL
	id, err := strconv.Atoi(categoryId) // Mengonversi ID dari string ke integer
	helper.PanicIfError(err)

	categoryUpdateRequest.Id = id 		// Menetapkan ID kategori yang akan diperbarui

	// 2. Melakukan validasi terhadap request

	// 3. Memanggil service untuk menyimpan kategori baru
	categoryResponse := controller.CategoryService.Update(request.Context(), categoryUpdateRequest)
	webResponse := web.WebResponse{
		Code:   http.StatusOK, 			// Kode status HTTP untuk OK
		Status: "OK",	 				// Deskripsi status HTTP
		Data:   categoryResponse,	 	// Data yang ingin dikirimkan dalam respons
	}

	// 4. Mengembalikan response yang sesuai dengan format JSON
	helper.WriteToResponse(writer, webResponse) 
}

func (controller *CategoryControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// Implementasi untuk menghapus kategori

	// 1. Mengambil ID kategori dari parameter URL
	categoryId := params.ByName("id") 	// Mengambil ID kategori dari parameter URL
	id, err := strconv.Atoi(categoryId) // Mengonversi ID dari string ke integer
	helper.PanicIfError(err)

	// 2. Memanggil service untuk menghapus kategori
	controller.CategoryService.Delete(request.Context(), id)

	// 3. Memanggil service untuk menghapus kategori 
	webResponse := web.WebResponse{
		Code:   http.StatusOK, 			// Kode status HTTP untuk OK
		Status: "OK",	 				// Deskripsi status HTTP
		// Data:   categoryResponse,	 	// Tidak ada Data yang ingin dikirimkan dalam respons
	}

	// 4. Mengembalikan response yang sesuai dengan format JSON
	helper.WriteToResponse(writer, webResponse) 
}

func (controller *CategoryControllerImpl) FindByid(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// Implementasi untuk menemukan kategori berdasarkan ID

	// 1. Mengambil ID kategori dari parameter URL
	categoryId := params.ByName("id") 	// Mengambil ID kategori dari parameter URL
	id, err := strconv.Atoi(categoryId) // Mengonversi ID dari string ke integer
	helper.PanicIfError(err)

	// 2. Memanggil service untuk menemukan kategori berdasarkan ID
	categoryResponse := controller.CategoryService.FindById(request.Context(), id)

	// 3. Memanggil service untuk menghapus kategori 
	webResponse := web.WebResponse{
		Code:   http.StatusOK, 			// Kode status HTTP untuk OK
		Status: "OK",	 				// Deskripsi status HTTP
		Data:   categoryResponse,	 	// Data yang ingin dikirimkan dalam respons
	}

	// 4. Mengembalikan response yang sesuai dengan format JSON
	helper.WriteToResponse(writer, webResponse)
}

func (controller *CategoryControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	// Implementasi untuk menemukan semua kategori
	
	// 1. Memanggil service untuk menemukan semua kategori
	categoriesResponses := controller.CategoryService.FindAll(request.Context())

	webResponse := web.WebResponse{
		Code:   http.StatusOK, 				// Kode status HTTP untuk OK
		Status: "OK",	 					// Deskripsi status HTTP
		Data:   categoriesResponses,	 	// Data yang ingin dikirimkan dalam respons
	}

	// 4. Mengembalikan response yang sesuai dengan format JSON
	helper.WriteToResponse(writer, webResponse) 
}