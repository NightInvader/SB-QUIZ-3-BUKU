package repository

import (
	"database/sql"
	"quiz/structs"
)

func GetAllBooks(db *sql.DB) (result []structs.Books, err error) {
	sql := "SELECT * FROM Books"

	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var book structs.Books

		err = rows.Scan(
			&book.ID,
			&book.Title,
			&book.Description,
			&book.ImageUrl,
			&book.ReleaseYear,
			&book.Price,
			&book.Price,
			&book.TotalPage,
			&book.Thickness,
			&book.CategoryID,
			&book.CreatedAt,
			&book.CreatedBy,
			&book.ModifiedAt,
			&book.ModifiedBy,
		)
		if err != nil {
			return nil, err
		}
		result = append(result, book)
	}
	return
}
