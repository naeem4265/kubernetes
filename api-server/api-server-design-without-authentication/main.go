package main

import (
	"github.com/gin-gonic/gin"

	"net/http"
)

type book struct {
	ID     string  `json: "id"`
	Title  string  `json: "title"`
	Writer string  `json: "writer"`
	Pages  int64   `json: "pages"`
	Price  float64 `json: "Price"`
	Count  int64   `json: "count"`
}

var albums = []book{
	{"1", "Programming in go", "X", 200, 1000.50, 100},
	{ID: "2", Title: "Programming in C++", Writer: "Y", Pages: 300, Price: 2000.50, Count: 200},
	{ID: "3", Title: "Programming in Java", Writer: "Z", Pages: 400, Price: 3000.50, Count: 300},
}

func getAlbums(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, albums)
}
func addAlbum(context *gin.Context) {
	var newAlbum book
	if err := context.BindJSON(&newAlbum); err != nil {
		return
	}
	albums = append(albums, newAlbum)
	context.IndentedJSON(http.StatusCreated, newAlbum)
}
func getAlbum(context *gin.Context) {
	id := context.Param("id")
	for _, x := range albums {
		if x.ID == id {
			context.IndentedJSON(http.StatusOK, x)
			return
		}
	}
	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
func updateById(context *gin.Context) {
	id := context.Param("id")
	var temp book
	if err := context.BindJSON(&temp); err != nil {
		return
	}
	for idx, _ := range albums {
		if albums[idx].ID == id {
			albums[idx] = temp
			context.IndentedJSON(http.StatusAccepted, temp)
			return
		}
	}
	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Id not found"})
}
func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", addAlbum)
	router.GET("/albums/:id", getAlbum)
	router.PUT("albums/:id", updateById)

	router.Run("localhost:8080")
}
