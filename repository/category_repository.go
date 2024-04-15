package repository

import (
	"context"
	"database/sql"
	"golang-rest-api/model/domain"
)

// CategoryRepository adalah interface yang mendefinisikan operasi-operasi CRUD
// (Create, Read, Update, Delete) yang berkaitan dengan entitas Category pada 
// penyimpanan data menggunakan database SQL.
type CategoryRepository interface {
	
	// Savectx bertugas untuk menyimpan atau membuat data entitas Category 
	// ke dalam penyimpanan data. Parameter ctx adalah konteks yang digunakan
	// untuk mengatur batasan waktu, pembatalan, dan nilai-nilai lain yang 
	// berhubungan dengan operasi tersebut. Parameter tx adalah transaksi SQL
	// yang sedang berjalan, dan parameter category adalah data Category yang akan disimpan.
	Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	
	// Update bertugas untuk mengupdate data entitas Category yang sudah ada 
	// di penyimpanan data. Parameter dan fungsinya sama dengan Savectx.
	Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	
	// Delete bertugas untuk menghapus data entitas Category dari penyimpanan data.
	// Parameter ctx adalah konteks, tx adalah transaksi SQL, dan category adalah 
	// data Category yang akan dihapus.
	Delete(ctx context.Context, tx *sql.Tx, category domain.Category)
	
	// FindById bertugas untuk mencari data entitas Category berdasarkan ID tertentu 
	// dari penyimpanan data. Parameter ctx adalah konteks, tx adalah transaksi SQL,
	// dan categoryId adalah ID dari Category yang akan dicari.
	FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error)
	
	// FindAll bertugas untuk mengambil semua data entitas Category dari penyimpanan data.
	// Parameter ctx adalah konteks dan tx adalah transaksi SQL.
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Category

	// context.Context digunakan sebagai parameter untuk setiap method dalam 
	// interface CategoryRepository untuk memastikan bahwa setiap operasi dapat 
	// diatur dan dikontrol dengan baik sesuai dengan kebutuhan aplikasi, 
	// seperti batasan waktu dan pembatalan operasi.
}
