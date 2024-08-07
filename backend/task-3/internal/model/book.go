package model

import "time"

type (
	Book struct {
		Id            string
		Title         string
		ISBN          string
		Author        string
		PublishedDate time.Time
		Genre         string
		Price         int64
		Quantity      int
		CreatedAt     time.Time
		UpdatedAt     time.Time
		DeletedAt     time.Time
	}

	Books []Book
)
