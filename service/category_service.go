package service

import (
	"context"
	"golang_restful_api/model/web"
)

// BUAT KONTRAKNYA DULU
type CategoryService interface {
	// Bisanya jumlah functionnya mengikuti jumlah endpoint/api-nya
	Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse
	Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse
	Delete(ctx context.Context, categoryId int)
	FindById(ctx context.Context, categoryId int) web.CategoryResponse
	FindAll(ctx context.Context) []web.CategoryResponse
}