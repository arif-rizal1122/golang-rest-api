package helper

// PanicIfError adalah sebuah fungsi utilitas yang digunakan untuk memeriksa 
// apakah terjadi error. Jika error tidak nil, maka fungsi akan memanggil 
// panic dengan error yang diterima sebagai argumen.
func PanicIfError(err error)  {
	// Memeriksa apakah error tidak nil.
	if err != nil {
		// Jika error tidak nil, panggil panic dengan error tersebut.
		panic(err)
	}
}
