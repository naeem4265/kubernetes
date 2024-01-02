package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/naeem4265/api-server/data"
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
			w.WriteHeader(http.StatusOK)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}
