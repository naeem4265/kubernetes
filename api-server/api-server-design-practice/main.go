package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type book struct {
	Id    string  `json: "id"`
	Title string  `json: "title"`
	Price float64 `json: "price"`
}

var albums = []book{
	{"1", "Programming in C", 12.23},
	{"2", "Programming in Go", 1256.32},
	{"3", "Programming in java", 12.3566},
}

func getAlbums(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, albums)
}
func postalbum(context *gin.Context) {
	var temp book
	if err := context.BindJSON(&temp); err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Data not found"})
		return
	}
	albums = append(albums, temp)
	context.IndentedJSON(http.StatusAccepted, temp)
}

func main() {
	fmt.Println("Server Created")
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postalbum)
	router.Run("localhost:8080")
}
