package news // логично, это модуль новостей)))

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

// функция, которую вызывает cервер HTTP
func GetAllNews(w http.ResponseWriter, r *http.Request) {
	// Разрешаем только GET
	if r.Method != http.MethodGet { // Делаем обязательно проверку! один handler = один метод (Защита от POST/PUT/DELETE)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	// Читаем просто данные
	posts, err := LoadPosts()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//Ответ клиенту
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

// HTTP-handler- Он принимает и отдает
func CreateNews(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var post Post

	err := json.NewDecoder(r.Body).Decode(&post) // r.Body — тело HTTP-запроса (там JSON)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// Передаем данные в storage (По факту просто вызывает бизнес логику)
	createdPost, err := CreatePost(post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Возвращает ответ клиенту
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdPost)
}
func DeleteNews(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// r.URL.Path = "/news/3"
	pathParts := strings.Split(r.URL.Path, "/")

	// ["", "news", "3"]
	if len(pathParts) != 3 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	idStr := pathParts[2]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = DeletePost(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func GetNewsByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// /news/3
	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) != 3 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(pathParts[2])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	posts, err := LoadPosts()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	for _, post := range posts {
		if post.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(post)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
}

func UpdateNews(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	pathParts := strings.Split(r.URL.Path, "/")
	if len(pathParts) != 3 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(pathParts[2])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var updatedPost Post
	err = json.NewDecoder(r.Body).Decode(&updatedPost)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	post, err := UpdatePost(id, updatedPost)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}
