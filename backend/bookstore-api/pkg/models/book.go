package models

type Book struct {
	ID       int    `json:"id"`
	NamaBuku string `json:"nama_buku"`
	Penerbit string `json:"penerbit"`
	Harga    int    `json:"harga_buku"`
}

func (Book) TableName() string {
	return "books"
}
