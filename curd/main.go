package main

import (
	"curd/model"
	"fmt"

	// "net/http"
	// "github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/db_crud?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err !=nil {
		fmt.Println("Koneksi Database GAGAL")
	}
	fmt.Println("koneksi berhasil")

// 	// dosen
	// insert
	// data := model.Dosen{
	// 	Nama: "Siti Nirmala Azwa N.",
	// 	Nidn: "20006020611",
	// 	Ttl: "2002/01/01",
	// 	Alamat: "REBAN TEBU",
	// }
// 	db.Create(&data)
// // update
// db.Model(&data).Where("id_dosen =?" , 1).Update("nidn", "200602061")

// delete
db.Where("id_dosen =?" , 3).Delete(&model.Dosen{})

// // read
// var dosen []model.Dosen
// db.Find(&dosen)
// for _, m :=range dosen{
// 	fmt.Println("Nama = "+m.Nama)
// 	fmt.Println("Nidn = "+m.Nidn)
// 	fmt.Println("Alamat= "+m.Alamat)
// 	fmt.Println("Ttl = "+m.Ttl)
// 	fmt.Println("==============================")
// }

// Tampilan data JSON
// r := gin.Default()
// 	r.GET("/getdosen", func(c *gin.Context) {
// 		var dosen []model.Dosen
// 		db.Find(&dosen)
// 		c.JSON(http.StatusOK, &dosen)
// 	})
// 	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	// var mhs model.Mahasiswa
	// 	mhs.Nama = "SITI"
	// 	mhs.Nim = "200602061"
	// 	mhs.Alamat = "RBT"
	
	// // berfungsi untuk menyimpan data ke database 
	// db.Create(&mhs)
	
	// // untuk update data pada file nama dengan id = 1
	// // db.Model(&mhs).Where("id= ?", 1).Update("nama", "siti")
	
	// // umtuk hapus data
	// // db.Where("id = ?", 2).Delete(&model.Mahasiswa{})
	
	// // Query atau mengambil semua database 
	//  var mahasiswa []model.Mahasiswa 
	//  db.Find(&mahasiswa)
	
	// for _, m := range mahasiswa{
	// 	fmt.Println("Nama = " + m.Nama)
	// 	fmt.Println("Nim = " + m.Nim)
	// 	fmt.Println("Alamat = " + m.Alamat)
	// 	fmt.Println("======================================")
	}


