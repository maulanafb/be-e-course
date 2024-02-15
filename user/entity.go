package user

import "time"

type User struct {
	ID        int
	Email     string
	Name      string
	Image     string
	Password  string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
