package web

type CategoryUpdateRequest struct {
	Id   int    `validate:"required" json:"id"`                 // Validasi untuk ID kategori
	Name string `validate:"required,min=1,max=200" json:"name"` // Validasi untuk nama kategori
}