package helper

import (
	"golang-rest-api/model/domain"
	"golang-rest-api/model/web"
)

// ToCategoryResponse adalah fungsi utilitas yang mengkonversi tipe data domain.Category 
// menjadi tipe data web.CategoryResponse.
func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	// Mengkonversi tipe data domain.Category menjadi tipe data web.CategoryResponse.
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

// ToCategoryResponses adalah fungsi utilitas yang mengkonversi slice dari tipe data domain.Category 
// menjadi slice dari tipe data web.CategoryResponse.
func ToCategoryResponses(categories []domain.Category) []web.CategoryResponse {
	// Membuat slice untuk menampung hasil konversi.
	var CategoryResponses []web.CategoryResponse

	// Iterasi melalui setiap elemen dalam slice categories.
	for _, category := range categories {
		// Memanggil fungsi ToCategoryResponse untuk setiap category dan menambahkan hasilnya 
		// ke dalam slice CategoryResponses.
		CategoryResponses = append(CategoryResponses, ToCategoryResponse(category))
	}

	// Mengembalikan slice yang berisi hasil konversi.
	return CategoryResponses
}
