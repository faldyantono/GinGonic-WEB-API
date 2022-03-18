package handler

import (
	"github.com/gin-gonic/gin"
	"time"
)

func QueryProduct(c *gin.Context) {
	name := c.Query("name")
	Id := c.Param("Id")
	harga := c.Query("harga")
	c.JSON(200, gin.H{
		"message":      "item name",
		"id":           Id,
		"Product name": name,
		"Price":        harga,
		"Time":         time.Now(),
	})
}
func GetItem(c *gin.Context) {
	name := c.Query("name")
	harga := c.Query("harga")
	c.JSON(200, gin.H{
		"message":      "item name",
		"product name": name,
		"Price":        harga,
		"Time":         time.Now(),
	})
}

//TUGAS NO 3 : bikin 5 router query string
func GetProduct1(c *gin.Context) {
	brand := c.Query("brand")
	c.JSON(200, gin.H{
		"message":    "Brand",
		"Nama Merek": brand,
	})
}
func GetProduct2(c *gin.Context) {
	class := c.Query("class")
	c.JSON(200, gin.H{
		"message":       "item class",
		"Product Class": class,
	})
}
func GetProduct3(c *gin.Context) {
	noId := c.Query("noId")
	c.JSON(200, gin.H{
		"message":    "Nomor id produk",
		"product Id": noId,
	})
}
func GetProduct4(c *gin.Context) {
	person := c.Query("person")
	c.JSON(200, gin.H{
		"message":   "product person",
		"person ke": person,
	})
}
func GetProduct5(c *gin.Context) {
	pet := c.Query("pet")
	c.JSON(200, gin.H{
		"message": "pet",
		"pet ke":  pet,
	})
}

//TUGAS NO 4 : bikin 5 router query param
func QueryParam1(c *gin.Context) {
	brand := c.Param("brand")
	c.JSON(200, gin.H{
		"message": "brand ke",
		"brand":   brand,
	})
}
func QueryParam2(c *gin.Context) {
	class := c.Param("class")
	c.JSON(200, gin.H{
		"message": "class item",
		"class":   class,
	})
}
func QueryParam3(c *gin.Context) {
	noId := c.Param("noId")
	c.JSON(200, gin.H{
		"message": "id ke",
		"Id":      noId,
	})
}
func QueryParam4(c *gin.Context) {
	person := c.Param("person")
	c.JSON(200, gin.H{
		"message": "person ke",
		"person":  person,
	})
}
func QueryParam5(c *gin.Context) {
	pet := c.Param("pet")
	c.JSON(200, gin.H{
		"message": "hello",
		"pet":     pet,
	})
}
