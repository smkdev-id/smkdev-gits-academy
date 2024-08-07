package models

type Book struct {
	ID         uint     `gorm:"primaryKey" json:"id"`
	Title      string   `json:"title"`
	Author     string   `json:"author"`
	IDKategori uint     `json:"id_kategori"`
	Kategori   Kategori `gorm:"foreignKey:IDKategori" json:"kategori"`
}

type Kategori struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Deskripsi string `json:"deskripsi"`
}
