package buku

type BukuRequest struct {
	Title string `json:"Title" binding:"required"`
	Price string `json:"Price" binding:"required"`
}
