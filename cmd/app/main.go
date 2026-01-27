package main

import (
	"log"
	"net/http"
	"news-crud/internal/news"
)

func main() {
	http.HandleFunc("/news/", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
		case http.MethodGet:
			news.GetAllNews(w, r)
		case http.MethodPost:
			news.CreateNews(w, r)
		case http.MethodDelete:
			news.DeleteNews(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
