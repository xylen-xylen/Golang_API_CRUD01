package api

import (
	"Golang_API_CRUD01/Internal/store"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type PostHandler struct {
	postStore store.PostStore
}

func NewPostHandler(ps store.PostStore) *PostHandler {
	return &PostHandler{
		postStore: ps,
	}
}
func (ph *PostHandler) HandlerGetPostById(w http.ResponseWriter, r *http.Request) {
	paramsPostId := chi.URLParam(r, "id")
	if paramsPostId == "" {
		http.NotFound(w, r)
		return
	}

	postId, err := strconv.ParseInt(paramsPostId, 10, 64)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "This is port id :%d\n", postId)
}
func (ph *PostHandler) HandlerCreatePost(w http.ResponseWriter, r *http.Request) {
	var post store.Post

	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, "Invalid request data", http.StatusBadRequest)
		return
	}

	createdPost, err := ph.postStore.CreatePost(&post)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "can't create post", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdPost)
}
