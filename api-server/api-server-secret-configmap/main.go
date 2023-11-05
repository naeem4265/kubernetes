package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt/v4"
	"github.com/naeem4265/api-server/handlers"
	"k8s.io/client-go/util/homedir"
	"log"
	"net/http"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// Creating our cluster config
	fmt.Println("Programm started")
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "Absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "Absolute path to the kubeconfig file")
	}
	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Clientset created")
	// Namespace where your Secret is located
	namespace := "default"
	secretName := "apiserver-secret" // Name of the Secret

	// Retrieve the Secret
	secret, err := clientset.CoreV1().Secrets(namespace).Get(context.TODO(), secretName, metav1.GetOptions{})
	if err != nil {
		panic(err.Error())
	}

	// Access the data in the Secret
	username, usernameExists := secret.Data["username"]
	password, passwordExists := secret.Data["password"]

	// Check if the environment variables are set
	if !usernameExists || !passwordExists {
		fmt.Println("Username or password not found in the Secret.")
		return
	}
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
