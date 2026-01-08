package main

import (
	"log"
	"net/http"

	"news-crud/internal/news"
)

func main() {
	// Регистрируем handler
	http.HandleFunc("/news", news.GetAllNews)

	log.Println("Server started on :8080")

	// Запускаем сервер
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
