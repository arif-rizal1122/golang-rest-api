package web


type CategoryUpdateRequest struct {
	// Id adalah identifikasi unik dari kategori yang akan diperbarui.
	Id int `validate:"required"`

	// Name adalah nama dari kategori yang akan diperbarui.
	// Nama kategori harus memiliki panjang minimal 1 karakter dan maksimal 200 karakter.
	Name string `validate:"required,max=200,min=1" json:"name"`
}
