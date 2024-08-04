package request

type (
	CreateBookRequest struct {
		Title             string `validate:"required,min=3,max=30" json:"title"`
		Description       string `validate:"required,min=5,max=500" json:"description"`
		Author            string `validate:"required,min=3,max=20" json:"author"`
		YearOfManufacture int    `validate:"required,gt=0" json:"year"`
		Stock             int    `validate:"required,gt=0" json:"stock"`
		Price             int    `validate:"required,gt=0" json:"price"`
	}

	UpdateBookRequest struct {
		Title             string `validate:"required,min=3,max=30" json:"title"`
		Description       string `validate:"required,min=5,max=500" json:"description"`
		Author            string `validate:"required,min=3,max=20" json:"author"`
		YearOfManufacture int    `validate:"required,gt=0" json:"year"`
		Stock             int    `validate:"required,gt=0" json:"stock"`
		Price             int    `validate:"required,gt=0" json:"price"`
		IsDisplayed       bool   `json:"is_displayed"`
	}

	SearchBookRequest struct {
		ISBN        string `json:"isbn"`
		Title       string `json:"title"`
		Year        int    `json:"year"`
		Author      string `json:"author"`
		Price       int    `json:"price"`
		IsDisplayed *bool  `json:"is_displayed"`
		StartDate   string `json:"start_date"`
		EndDate     string `json:"end_date"`
	}
)
