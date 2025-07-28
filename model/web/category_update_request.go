package web

type CategoryUpdateRequest struct {
	Id   int    `validate:"required"`               // Validasi untuk ID kategori
	Name string `validate:"required,min=1,max=200"` // Validasi untuk nama kategori
}