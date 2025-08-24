package structs

import "time"

type Books struct {
	ID          int    `gorm:"primarykey"`
	Title       string `gorm:"type:varchar(255);not null"`
	Description string
	ImageUrl    string
	ReleaseYear int
	Price       int
	TotalPage   int
	Thickness   string
	CategoryID  int
	Category    Category `gorm:"foreignKey:CategoryID"`
	CreatedAt   time.Time
	CreatedBy   string
	ModifiedAt  time.Time
	ModifiedBy  time.Time
}

type Category struct {
	ID         int    `gorm:"primaryKey"`
	Name       string `gorm:"type:varchar(255);not null"`
	CreatedAt  time.Time
	CreatedBy  string
	ModifiedAt time.Time
	ModifiedBy time.Time
	Books      []Books
}

type User struct {
	ID         int    `gorm:"primaryKey"`
	Username   string `gorm:"uniqueIndex;not null"`
	Password   string `gorm:"not null"`
	CreatedAt  time.Time
	CreatedBy  string
	ModifiedAt time.Time
	ModifiedBy time.Time
}
