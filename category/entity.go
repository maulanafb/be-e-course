package category

import "time"

type Category struct {
	ID        string
	Title     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
