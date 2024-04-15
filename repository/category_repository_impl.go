package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang-rest-api/helper"
	"golang-rest-api/model/domain"
)

// CategoryRepositoryImpl adalah implementasi dari interface CategoryRepository.
// Ini adalah struktur yang akan digunakan untuk mengimplementasikan fungsi-fungsi CRUD
// untuk entitas Category pada penyimpanan data menggunakan database SQL.
type CategoryRepositoryImpl struct {
	// Anda dapat menambahkan field tambahan di sini jika diperlukan

}



func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}




// Save adalah method untuk menyimpan atau membuat data entitas Category 
// ke dalam penyimpanan data. 
func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	// SQL statement untuk insert data ke dalam tabel category dengan parameter name.
	SQL := "insert into category(name) values (?)"
	// Eksekusi SQL dengan transaksi dan konteks yang diberikan.
	result, err := tx.ExecContext(ctx, SQL, category.Name)
	helper.PanicIfError(err)

	// Mendapatkan ID dari data yang baru saja disimpan.
	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	// Mengatur ID yang baru didapatkan ke dalam struct category.
	category.Id = int(id)
	return category
}

// Update adalah method untuk mengupdate data entitas Category yang sudah ada 
// di penyimpanan data. 
func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	// SQL statement untuk update data pada tabel category berdasarkan ID dengan parameter name.
	SQL := "update category set name = ? where id = ?"
	// Eksekusi SQL dengan transaksi, konteks, dan parameter yang diberikan.
	_ , err := tx.ExecContext(ctx, SQL, category.Name, category.Id)
	helper.PanicIfError(err)

	return category
}

// Delete adalah method untuk menghapus data entitas Category dari penyimpanan data.
func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category)  {
	// SQL statement untuk menghapus data dari tabel category berdasarkan ID.
	SQL := "delete from category where id = ?"
	// Eksekusi SQL dengan transaksi, konteks, dan parameter yang diberikan.
	_ , err := tx.ExecContext(ctx, SQL, category.Id)
	helper.PanicIfError(err)
}

// FindById adalah method untuk mencari data entitas Category berdasarkan ID tertentu 
// dari penyimpanan data.
func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) {
	// SQL statement untuk mengambil data dari tabel category berdasarkan ID.
	SQL := "select id, name from category where id = ?"
	// Eksekusi SQL dengan transaksi, konteks, dan parameter yang diberikan.
	rows, err := tx.QueryContext(ctx, SQL, categoryId)
	helper.PanicIfError(err)

	defer rows.Close()

	category := domain.Category{}
	// Jika data ditemukan, isi struct category dengan data dari baris tersebut.
	if rows.Next() {
		rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		return category, nil
	} else {
		// Jika data tidak ditemukan, kembalikan error "category is not found".
		return category, errors.New("category is not found")
	}
}

// FindAll adalah method untuk mengambil semua data entitas Category dari penyimpanan data.
func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	// SQL statement untuk mengambil semua data dari tabel category.
	SQL := "select id, name from category"
	// Eksekusi SQL dengan transaksi dan konteks yang diberikan.
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)

	defer rows.Close()

	var categories []domain.Category
	// Loop melalui setiap baris hasil query dan tambahkan ke slice categories.
	for rows.Next() {
		category := domain.Category{}
		rows.Scan(&category.Id, &category.Name)
		categories = append(categories, category)
	}

	return categories
}
