package entity

// Book Struct
type Book struct {
	ID       int32   `json:"book_id" db:"book_id"`
	Title    string  `json:"book_title" db:"book_name"`
	Author   string  `json:"book_author_name" db:"book_author_name"`
	Category string  `json:"book_category" db:"book_category"`
	Rate     float32 `json:"book_rate" db:"book_category"`
}
