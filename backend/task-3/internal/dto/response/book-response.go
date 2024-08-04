package response

type (
	BookResponse struct {
		Id                string `json:"id"`
		Isbn              string `json:"isbn"`
		Title             string `json:"book_title"`
		Description       string `json:"book_description"`
		Author            string `json:"author"`
		YearOfManufacture int    `json:"year_of_manufacture"`
		Stock             int    `json:"stock"`
		Price             int    `json:"price"`
		IsDisplayed       bool   `json:"is_displayed"`
		CreatedAt         string `json:"created_at"`
		UpdatedAt         string `json:"updated_at"`
	}
)
