package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"webapi/buku"
	"webapi/handler"
)

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:@tcp(127.0.0.1:3306)/dbbuku?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Db Connection Error")
	}
	db.AutoMigrate(&buku.Buku{})
	//memanggil repository
	bukuRepository := buku.NewRepository(db)
	buku := buku.Buku{
		Title:       "$100 Startup",
		Description: "Good Book",
		Price:       200000,
		Rating:      4,
		Discount:    0,
	}
	bukuRepository.Create(buku)

	//fmt.Println("Database connected")
	//CRUD(Create Read Update Delete)
	//first part (Create)
	//buku := buku.Buku{}
	//buku.Title = "Top Gear"
	//buku.Price = 120000
	//buku.Discount = 20
	//buku.Rating = 5
	//buku.Description = "All of Top Gear 2012 Season is all sumarized here!."
	//err = db.Create(&buku).Error
	//if err != nil {
	//	fmt.Println("===========================")
	//	fmt.Println("Error creating book record")
	//	fmt.Println("===========================")
	//}
	//var buku buku.Buku
	//Second Part of CRUD, Read
	//err = db.Debug().Where("Id = ?", 2).First(&buku).Error
	//if err != nil {
	//	fmt.Println("===========================")
	//	fmt.Println("Error finding book record")
	//	fmt.Println("===========================")
	//}
	//for _, b := range buku {
	//	fmt.Println("Title :", b.Title)
	//	fmt.Println("Buku object %v", b)
	//}
	//Third Part of CRUD, Update
	//buku.Title = "Top Gear(Revised Edition)"
	//err = db.Save(&buku).Error
	//if err != nil {
	//	fmt.Println("Updating Book Record")
	//}
	//Fourth Part of CRUD (DELETE)
	//err = db.Delete(&buku).Error
	//if err != nil {
	//	fmt.Println("=============================")
	//	fmt.Println("Error Deleting Book Record")
	//	fmt.Println("=============================")
	//}
	//err = db.Create(&buku).Error
	//if err != nil {
	//	fmt.Println("Error membuat data")
	//}
	//
	//err = db.Find(&buku).Error
	//if err != nil {
	//	fmt.Println("Error membuat data")
	//}
	//err = db.Debug().Where("id = ?", 1).Find(&buku).Error
	//if err != nil {
	//	fmt.Println("Error menermukan data")
	//}
	//buku.Title = "Hati hati di jalan"
	//db.Save(&buku)
	//if err != nil {
	//	fmt.Println("Error update data data")
	//}
	//err = db.Delete(&buku).Error
	//if err != nil {
	//	fmt.Println("Error update data data")
	//}
	//fmt.Println("Title", buku.Title)
	//fmt.Println("Deskripsi", buku.Description)

	r := gin.Default()
	v1 := r.Group("/v1")
	//TUGAS NO 1. : localhost:8080/item?judul=filosofi-teras&harga=76000
	v1.GET("/item", handler.GetItem)
	//TUGAS NO 2 : localhost:8080/id/2/produk?name=tv&harga=4000000
	v1.GET("id/:Id/produk", handler.QueryProduct)
	//TUGAS NO 3 : bikin 5 router query string
	v1.GET("/item1", handler.GetProduct1)
	v1.GET("/item2", handler.GetProduct2)
	v1.GET("/item3", handler.GetProduct3)
	v1.GET("/item4", handler.GetProduct4)
	v1.GET("/item5", handler.GetProduct5)
	//TUGAS NO 4 : bikin 5 query param
	v1.GET("/product1/:brand", handler.QueryParam1)
	v1.GET("/product2/:class", handler.QueryParam2)
	v1.GET("/product3/:noId", handler.QueryParam3)
	v1.GET("/product4/:person", handler.QueryParam4)
	v1.GET("/product5/:pet", handler.QueryParam5)
	//TUGAS POST
	v1.POST("/mahasiswa", handler.MahasiswaHandl)
	r.Run() // listen and serve on 0.0.0.0:8080
}
