package main

import (
	"log"
	"net/http"
	"news-crud/internal/news"
	"strings"
)

func main() {
	// Раздельные обработчики для чистоты
	http.HandleFunc("/news", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Запрос к /news: %s", r.Method)
		switch r.Method {
		case http.MethodGet:
			news.GetAllNews(w, r)
		case http.MethodPost:
			news.CreateNews(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/news/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Запрос к /news/: %s %s", r.Method, r.URL.Path)

		// Проверяем, что после /news/ есть ID
		path := strings.TrimPrefix(r.URL.Path, "/news/")
		if path == "" {
			http.Error(w, "ID is required", http.StatusBadRequest)
			return
		}

		switch r.Method {
		case http.MethodGet:
			news.GetNewsByID(w, r)
		case http.MethodPut:
			news.UpdateNews(w, r)
		case http.MethodDelete:
			news.DeleteNews(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	// Простой обработчик для корня
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("Сервер API новостей\nИспользуй:\nGET /news - все новости\nPOST /news - создать\nGET /news/1 - получить по ID"))
	})

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
