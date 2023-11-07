package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/naeem4265/api-server/data"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

func GetAlbums(w http.ResponseWriter, r *http.Request) {
	albumJSON, err := json.Marshal(data.Albums)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Write(albumJSON)
	w.WriteHeader(http.StatusOK)
}

func PostAlbum(w http.ResponseWriter, r *http.Request) {
	var temp data.Book
	if err := json.NewDecoder(r.Body).Decode(&temp); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	data.Albums = append(data.Albums, temp)

	// First, marshal the JSON data
	jsonData, err := json.Marshal(data.Albums)
	filePath := "book/book.txt"
	dirPath := filepath.Dir(filePath)
	err = os.MkdirAll(dirPath, os.ModePerm)
	err = ioutil.WriteFile(filePath, jsonData, 0777)
	if err != nil {
		fmt.Println(err)
	}

	w.WriteHeader(http.StatusOK)
}

func GetAlbumById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	for idx, _ := range data.Albums {
		a := data.Albums[idx]
		if a.Id == id {
			albumJSON, err := json.Marshal(data.Albums[idx])
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			w.Write(albumJSON)
			w.WriteHeader(http.StatusOK)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func PutAlbum(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	for idx, _ := range data.Albums {
		a := data.Albums[idx]
		if a.Id == id {
			var temp data.Book
			if err := json.NewDecoder(r.Body).Decode(&temp); err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			data.Albums[idx] = temp

			// First, marshal the JSON data
			jsonData, err := json.Marshal(data.Albums)
			filePath := "book/book.txt"
			dirPath := filepath.Dir(filePath)
			err = os.MkdirAll(dirPath, os.ModePerm)
			err = ioutil.WriteFile(filePath, jsonData, 0777)
			if err != nil {
				fmt.Println(err)
			}

			w.WriteHeader(http.StatusOK)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func DeleteAlbum(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	for idx, _ := range data.Albums {
		a := data.Albums[idx]
		if a.Id == id {
			data.Albums = append(data.Albums[:idx], data.Albums[idx+1:]...)

			// First, marshal the JSON data
			jsonData, err := json.Marshal(data.Albums)
			filePath := "book/book.txt"
			dirPath := filepath.Dir(filePath)
			err = os.MkdirAll(dirPath, os.ModePerm)
			err = ioutil.WriteFile(filePath, jsonData, 0777)
			if err != nil {
				fmt.Println(err)
			}

			w.WriteHeader(http.StatusOK)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}
