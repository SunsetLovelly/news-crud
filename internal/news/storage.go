package news

import (
	"encoding/json"
	"os"
)

const dataFile = "data.json"

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
