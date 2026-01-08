package news // логично, это модуль новостей)))

import (
	"encoding/json"
	"net/http"
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
