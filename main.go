package main

import (
	"github.com/gin-gonic/gin"
	"webapi/handler"
)

func main() {
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
