package category

import "time"

type Category struct {
	ID        int
	Title     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
