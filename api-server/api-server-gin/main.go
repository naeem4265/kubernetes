package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = string("my_secret_key")

var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

// Create a struct that models the structure of a user, both in the request body, and in the DB
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func Signin(context *gin.Context) {
	var creds Credentials

	err := context.BindJSON(&creds)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Error"})
		return
	}

	expectedPassword, ok := users[creds.Username]
	if !ok || expectedPassword != creds.Password {
		context.IndentedJSON(http.StatusUnauthorized, creds)
		return
	}

	context.SetCookie(creds.Username, jwtKey, 50000, "/", "localhost", false, false)
	context.String(http.StatusOK, "Cookie has been set")
	/*
		expirationTime := time.Now().Add(5 * time.Minute)
		claims := &Claims{
			Username: creds.Username,
			RegisteredClaims: jwt.RegisteredClaims{
				// In JWT, the expiry time is expressed as unix milliseconds
				ExpiresAt: jwt.NewNumericDate(expirationTime),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(jwtKey)
		if err != nil {
			context.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Try agin"})
			return
		}
	*/

}

type book struct {
	Id    string  `json: "id"`
	Title string  `json: "title"`
	Price float64 `json: "price"`
}

var Albums = []book{
	{"1", "Programming in C", 1000},
	{"2", "Programming in Java", 2000},
	{"3", "Programming in Go", 3000},
}

func getAlbums(context *gin.Context) {
	cookie, err := context.Cookie("user1")
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "cookie not found"})
		return
	}
	context.IndentedJSON(http.StatusOK, cookie)
	context.IndentedJSON(http.StatusOK, Albums)
	return

}

func main() {
	router := gin.Default()
	router.POST("/signin", Signin)
	router.GET("/albums", getAlbums)

	router.Run("localhost:8080")
}
