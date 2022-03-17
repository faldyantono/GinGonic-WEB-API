package main

import (
	"github.com/gin-gonic/gin"
	"webapi/handler"
)

func main() {
	r := gin.Default()
	//TUGAS NO 1. : localhost:8080/item?judul=filosofi-teras&harga=76000
	r.GET("/item", handler.GetItem)
	//TUGAS NO 2 : localhost:8080/id/2/produk?name=tv&harga=4000000
	r.GET("produk/:id/produk", handler.QueryProduct)
	//TUGAS NO 3 : bikin 5 router query string
	r.GET("/item1", handler.GetProduct1)
	r.GET("/item2", handler.GetProduct2)
	r.GET("/item3", handler.GetProduct3)
	r.GET("/item4", handler.GetProduct4)
	r.GET("/item5", handler.GetProduct5)
	//TUGAS NO 4 : bikin 5 query param
	r.GET("/product1/:brand", handler.QueryParam1)
	r.GET("/product2/:class", handler.QueryParam2)
	r.GET("/product3/:noId", handler.QueryParam3)
	r.GET("/product4/:person", handler.QueryParam4)
	r.GET("/product5/:pet", handler.QueryParam5)
	r.Run() // listen and serve on 0.0.0.0:8080
}
