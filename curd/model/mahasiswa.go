package model

// struk model untuk mempersentasikan tabel database
type Mahasiswa struct {
	ID     int
	Nama   string
	Nim    string
	Alamat string
}

// berfungsi untuk menyimpan data ke database 
// db.Create(&mhs)

// untuk update data pada file nama dengan id = 1
// db.Model(&mhs).Where("id= ?", 1).Update("nama", "siti")

// umtuk hapus data
// db.Where("id = ?", 2).Delete(&model.Mahasiswa{})
