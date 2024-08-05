package helper

type Request struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Genre       string    `json:"genre"`
	ISBN        string    `json:"isbn"`
	Author      string    `json:"author"`
	Price       float64   `json:"price"`
}