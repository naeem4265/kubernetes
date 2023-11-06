package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt/v4"
	"github.com/naeem4265/api-server/handlers"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// Access the data in the Secret

	fmt.Println("Programm started")
	listenport, err := ioutil.ReadFile("/config/listenPort.config")
	if err != nil {
		fmt.Printf("Error reading config file: %v\n", err)
		return
	}
	fmt.Printf("Port: %s\n", string(listenport))

	usersDir := "/users" // The directory containing the files

	files, err := ioutil.ReadDir(usersDir)
	if err != nil {
		log.Fatalf("Error reading directory: %v", err)
	}

	for _, file := range files {
		if !file.IsDir() {
			filePath := usersDir + "/" + file.Name()
			fileContent, err := ioutil.ReadFile(filePath)
			if err != nil {
				//log.Printf("Error reading file %s: %v", filePath, err)
				continue
			} else {
				fmt.Printf("File: %s\nContent: %s\n", file.Name(), string(fileContent))
				handlers.MapUsernamePassword(file.Name(), string(fileContent))
			}
		}
	}

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

	fmt.Printf("Server started at :%s", string(listenport))
	//portStr := string(listenport)
	//port := ":" + portStr
	//log.Fatal(http.ListenAndServe(port, router))

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
