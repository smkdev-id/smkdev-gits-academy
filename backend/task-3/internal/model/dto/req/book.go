package req

import "time"

type (
	CreateRequest struct {
		Id            string    `json:"id"`
		Title         string    `json:"title"`
		ISBN          string    `json:"isbn"`
		Author        string    `json:"author"`
		PublishedDate time.Time `json:"published_date"`
		Genre         string    `json:"genre"`
		Price         int64     `json:"price"`
		Quantity      int       `json:"quantity"`
		CreatedAt     time.Time `json:"created_at"`
		UpdatedAt     time.Time `json:"updated_at"`
		DeletedAt     time.Time `json:"deleted_at"`
	}

	UpdateRequest struct {
		Id            string    `json:"id" validate:"required"`
		Title         string    `json:"title" validate:"omitempty"`
		ISBN          string    `json:"isbn" validate:"omitempty"`
		Author        string    `json:"author" validate:"omitempty"`
		PublishedDate time.Time `json:"published_date" validate:"omitempty"`
		Genre         string    `json:"genre" validate:"omitempty"`
		Price         int64     `json:"price" validate:"omitempty"`
		Quantity      int       `json:"quantity" validate:"omitempty"`
		UpdatedAt     time.Time `json:"updated_at"`
	}
	
	DeleteRequest struct {
		Id string `json:"id" validate:"required"`
	}
)
