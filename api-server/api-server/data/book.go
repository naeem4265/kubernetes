package data

type Book struct {
	Id    string  `json:"id"`
	Title string  `json:"title,omitempty"`
	Price float64 `json:"price"`
}

var Albums = []Book{
	{"1", "Programming in C", 1000},
	{"2", "Programming in Java", 2000},
	{"3", "Programming in Go", 3000},
}
