package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"time"
)

//TUGAS MEMBUAT POST
type Mahasiswainput struct {
	ID     string  `json:"ID_Mahasiswa" binding:"required"`
	Nama     string  `json:"Nama_Mahasiswa" binding:"required"`
	Prodi    string  `json:"Program_Studi" binding:"required"`
	Fakultas     int     `json:"Fakultas" binding:"required"`
	NIM      int `json:"NIM" binding:"required"`
	Angkatan int  `json:"Tahun_Angkatan" binding:"required"`
}

func MahasiswaHandl(c *gin.Context) {
	//nama mahasiswa
	var mahasiswaInput Mahasiswainput

	err := c.ShouldBindJSON(&mahasiswaInput)
	//jika if tidak sama dengan nil(error), maka deklarasikan variabel array
	if err != nil {
		errorMessages := []string{}
		//dimana kita akan melakukan perulanga n for untuk error, sehingga eror akan memprint pesan error didalam variabel errorMessage
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error %s, message: %s", e.Field(), e.ActualTag())
			//selanjutnya kita append errorMessage di dalam variabel array errorMessages sehingga message error-
			//dapat dideklarasikan lebih dari satu
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message"		:        "Berhasil input data",
		"ID"			: mahasiswaInput.ID,
		"Nama Mahasiswa": mahasiswaInput.Nama,
		"Program Studi"	:  mahasiswaInput.Prodi,
		"Fakultas"		: mahasiswaInput.Fakultas,
		"NIM "			:  mahasiswaInput.NIM,
		"Angkatan "		:  mahasiswaInput.Angkatan,
		"Time"			:	           time.Now(),
	})
}
