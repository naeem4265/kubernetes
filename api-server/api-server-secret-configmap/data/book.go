package data

type Book struct {
	Id    string  `json:"id"`
	Title string  `json:"title,omitempty"`
	Price float64 `json:"price"`
}

var Albums []Book
