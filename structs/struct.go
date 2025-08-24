package structs

import "time"

type Books struct {
	ID          int
	Title       string
	Description string
	ImageUrl    string
	ReleaseYear int
	Price       int
	TotalPage   int
	Thickness   string
	CategoryID  Category
	CreatedAt   time.Time
	CreatedBy   string
	ModifiedAt  time.Time
	ModifiedBy  time.Time
}

type Category struct {
	ID         int
	Name       string
	CreatedAt  time.Time
	CreatedBy  string
	ModifiedAt time.Time
	ModifiedBy time.Time
}

type User struct {
	ID         int
	Username   string
	Password   string
	CreatedAt  time.Time
	CreatedBy  string
	ModifiedAt time.Time
	ModifiedBy time.Time
}
