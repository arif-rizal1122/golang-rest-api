package service

import (
	"context"
	"golang-rest-api/model/web"
)

// CategoryService adalah interface yang mendefinisikan kontrak untuk operasi
// yang dapat dilakukan pada entitas Category. Interface ini digunakan untuk
// mengabstraksi logika bisnis dari operasi CRUD yang akan diimplementasikan
// oleh CategoryServiceImpl atau service lainnya.
type CategoryService interface {
	// Create adalah method untuk membuat kategori baru.
	Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse

	// Update adalah method untuk memperbarui kategori berdasarkan ID.
	Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse

	// Delete adalah method untuk menghapus kategori berdasarkan ID.
	Delete(ctx context.Context, categoryId int)

	// FindById adalah method untuk mencari kategori berdasarkan ID.
	FindById(ctx context.Context, categoryId int) web.CategoryResponse

	// FindAll adalah method untuk mencari semua kategori.
	FindAll(ctx context.Context) []web.CategoryResponse
}
