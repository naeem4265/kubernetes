package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt/v4"
	"github.com/naeem4265/api-server/handlers"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	// Access the data in the Secret
	username := os.Getenv("APISERVER_USERNAME")
	password := os.Getenv("APISERVER_PASSWORD")
	if username == "" || password == "" {
		fmt.Printf("APISERVER_USERNAME or APISERVER_PASSWORD environment variables are not set.\n")
		return
	}
	listenport, err := ioutil.ReadFile("/config/listenPort.config")
	if err != nil {
		fmt.Printf("Error reading config file: %v\n", err)
		return
	}
	// Now you can use 'fileContent' in your application.
	fmt.Printf("Config Content: %s\n", string(listenport))

	// Now you can use the username and password in your Go program
	fmt.Printf("Username: %s\n", username)
	fmt.Printf("Password: %s\n", password)
	handlers.MapUsernamePassword(string(username), string(password))

	fmt.Println("Creating api-server")
	router := chi.NewRouter()

	router.Post("/signin", handlers.SignIn)
	router.Get("/signout", handlers.SignOut)

	router.Route("/albums", func(r chi.Router) {
		r.Use(middleware)
		r.Get("/", handlers.GetAlbums)
		r.Get("/{id}", handlers.GetAlbumById)
		r.Put("/{id}", handlers.PutAlbum)
		r.Post("/", handlers.PostAlbum)
		r.Delete("/{id}", handlers.DeleteAlbum)
	})

	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check for the "token" cookie
		c, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			// For any other type of error, return a bad request status
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		tknStr := c.Value

		claims := &handlers.Claims{}
		tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
			return handlers.JWTKey, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				// Token signature is invalid, return unauthorized status
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			// For any other error while parsing claims, return a bad request status
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if !tkn.Valid {
			// Token is not valid, return unauthorized status
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// If token is valid, continue to the next handler
		next.ServeHTTP(w, r)
	})
}
