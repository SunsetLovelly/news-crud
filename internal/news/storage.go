package news

import (
	"encoding/json"
	"os"
)

const dataFile = "internal/news/data.json"

func LoadPosts() ([]Post, error) { //чтение файла
	data, err := os.ReadFile(dataFile)
	if err != nil {
		if os.IsNotExist(err) {
			return []Post{}, nil
		}
		return nil, err
	}

	var posts []Post
	if len(data) == 0 {
		return []Post{}, nil
	}

	err = json.Unmarshal(data, &posts)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func SavePosts(posts []Post) error { //запись (мы делаем права доступа и делаем красиво через MarshalIndent
	data, err := json.MarshalIndent(posts, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(dataFile, data, 0644)
}

// Загружаем все посты из json.
func CreatePost(post Post) (Post, error) {
	posts, err := LoadPosts() // если ошибка сразу выходим!
	if err != nil {
		return Post{}, err // пустой объект, чтобы не возвращать мусор
	}

	// Генерируем ID, Ищем максимальный Id, Cохраняем обратно. (ручная эмуляция базы данных)
	maxID := 0
	for _, p := range posts {
		if p.ID > maxID {
			maxID = p.ID
		}
	}
	post.ID = maxID + 1 // присваиваем каждому свой уникальный ID

	posts = append(posts, post)

	err = SavePosts(posts)
	if err != nil {
		return Post{}, err
	}

	return post, nil
}

// Добавляем функцию удаления по Id
func DeletePost(id int) error {
	posts, err := LoadPosts()
	if err != nil {
		return err
	}

	// Новый массив без удаляемого поста
	newPosts := make([]Post, 0)

	for _, post := range posts {
		if post.ID != id {
			newPosts = append(newPosts, post)
		}
	}

	// Сохраняем обновлённый список
	return SavePosts(newPosts)
}
