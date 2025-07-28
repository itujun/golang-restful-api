package service

import (
	"context"
	"database/sql"
	"golang_restful_api/helper"
	"golang_restful_api/model/domain"
	"golang_restful_api/model/web"
	"golang_restful_api/repository"

	"github.com/go-playground/validator"
)

type CategoryServiceImpl struct {
	CategoryRepository 	repository.CategoryRepository // Menyatakan bahwa CategoryServiceImpl adalah implementasi dari CategoryRepository
	DB                 	*sql.DB                       // Menyatakan bahwa CategoryServiceImpl memiliki koneksi ke database
	Validate			*validator.Validate
}

func (service *CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	// Lakukan validasi terhadap request sebelum melakukan transaksi
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	// Membuat transaksi baru
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// Mwmbuat objek kategori dari request
	category := domain.Category{
		Name: request.Name,
	}

	// Memanggil repository Save untuk menyimpan kategori baru
	category = service.CategoryRepository.Save(ctx, tx, category) 

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	// Lakukan validasi terhadap request sebelum melakukan transaksi
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	
	// Membuat transaksi baru
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// Mencari category yang hendak diupdate
	category, err := service.CategoryRepository.FindById(ctx, tx, request.Id)
	helper.PanicIfError(err)

	// Jika kategori sudah ketemu, lakukan pembaruan data kategori
	category.Name = request.Name

	// Memanggil repository Update untuk memperbarui kategori
	category = service.CategoryRepository.Update(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Delete(ctx context.Context, categoryId int) {
	// Membuat transaksi baru
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// Mencari category yang hendak dihapus
	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	helper.PanicIfError(err)

	// Memanggil repository Delete untuk menghapus kategori
	service.CategoryRepository.Delete(ctx, tx, category)

	// Tidak ada return value karena ini adalah operasi void
}

func (service *CategoryServiceImpl) FindById(ctx context.Context, categoryId int) web.CategoryResponse {
	// Membuat transaksi baru
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// Mencari category berdasarkan ID
	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	helper.PanicIfError(err)

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {
	// Membuat transaksi baru
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// Mencari semua kategori
	categories := service.CategoryRepository.FindAll(ctx, tx)

	return helper.ToCategoryResponses(categories)
}