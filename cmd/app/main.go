package main

import (
	"log"
	"net/http"

	"news-crud/internal/news"
)

func main() {
	// Регистрируем handler
	http.HandleFunc("/news", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			news.GetAllNews(w, r)
		case http.MethodPost:
			news.CreateNews(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	// Логируем запуск сервера
	log.Println("Server started on :8080")

	// Запускаем HTTP сервер (блокирует выполнение)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
