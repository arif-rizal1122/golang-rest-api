package service

import (
	"context"
	"database/sql"
	"golang-rest-api/exception"
	"golang-rest-api/helper"
	"golang-rest-api/model/domain"
	"golang-rest-api/model/web"
	"golang-rest-api/repository"

	"github.com/go-playground/validator/v10"
)

// CategoryServiceImpl adalah implementasi dari service.CategoryService.
// Struktur ini memiliki akses ke CategoryRepository untuk manipulasi data kategori
// dan akses ke database (*sql.DB) untuk operasi database.
type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}


func NewCategoryService(categoryRepository repository.CategoryRepository, DB *sql.DB, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
		DB: DB,
		Validate: validate,
	}
}



// Create adalah method untuk membuat kategori baru.
func (service *CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	// Validasi dulu  request nya
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	// Mulai transaksi database.
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	// Menangani commit atau rollback transaksi.
	defer helper.CommitOrRollback(tx)

	// Membuat objek domain.Category dari request.
	category := domain.Category{
		Name: request.Name,
	}

	// Menyimpan kategori ke database.
	category = service.CategoryRepository.Save(ctx, tx, category)
	return helper.ToCategoryResponse(category)
}

// Update adalah method untuk memperbarui kategori berdasarkan ID.
func (service *CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	// Validasi dulu  request nya
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	// Mulai transaksi database.
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	// Menangani commit atau rollback transaksi.
	defer helper.CommitOrRollback(tx)

	// Mencari kategori berdasarkan ID.
	category, err := service.CategoryRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	// Memperbarui nama kategori.
	category.Name = request.Name

	// Menyimpan perubahan ke database.
	category = service.CategoryRepository.Update(ctx, tx, category)
	return helper.ToCategoryResponse(category)
}

// Delete adalah method untuk menghapus kategori berdasarkan ID.
func (service *CategoryServiceImpl) Delete(ctx context.Context, categoryId int) {
	// Mulai transaksi database.
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// Mencari kategori berdasarkan ID.
	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	// Menghapus kategori dari database.
	service.CategoryRepository.Delete(ctx, tx, category)
}

// FindById adalah method untuk mencari kategori berdasarkan ID.
func (service *CategoryServiceImpl) FindById(ctx context.Context, categoryId int) web.CategoryResponse {
	// Mulai transaksi database.
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// Mencari kategori berdasarkan ID.
	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToCategoryResponse(category)
}

// FindAll adalah method untuk mencari semua kategori.
func (service *CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {
	// Mulai transaksi database.
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// Mencari semua kategori dari database.
	categories := service.CategoryRepository.FindAll(ctx, tx)
	return helper.ToCategoryResponses(categories)
}
