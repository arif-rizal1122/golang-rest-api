package helper

import "database/sql"

// CommitOrRollback adalah fungsi utilitas yang digunakan untuk menyelesaikan transaksi database.
// Jika terjadi panic (biasanya karena error), transaksi akan dirollback. 
// Jika tidak ada panic, transaksi akan di-commit.
func CommitOrRollback(tx *sql.Tx) {
	// Menggunakan recover() untuk menangkap panic yang mungkin terjadi.
	err := recover()

	// Memeriksa apakah ada error yang tertangkap.
	if err != nil {
		// Jika terjadi panic, melakukan rollback transaksi.
		errorRollBack := tx.Rollback()
		PanicIfError(errorRollBack)
		// Melempar kembali error yang diterima untuk mengetahui penyebab panic.
		panic(err)
	} else {
		// Jika tidak ada panic, melakukan commit transaksi.
		errorCommit := tx.Commit()
		PanicIfError(errorCommit)
	}
}
