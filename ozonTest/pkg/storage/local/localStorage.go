package local

import (
	"fmt"
	"ozonTest/pkg/models"
)

type LocalStorage struct {
	Long  map[string]string
	Short map[string]string
}

func NewLocalStorage() *LocalStorage {
	long := make(map[string]string)
	short := make(map[string]string)
	return &LocalStorage{Long: long, Short: short}
}

func (storage LocalStorage) CreateShortLink(link models.Link) (string, error) {
	fmt.Println("Защли")
	if _, ok := storage.Short[link.Long]; ok {
		return storage.Short[link.Long], nil
	}

	storage.Short[link.Long] = link.Short

	if _, ok := storage.Long[link.Short]; ok {
		var newLink models.Link
		for {
			newLink = models.NewLink(link.Long)
			if newLink.Short != link.Short {
				break
			}
		}
		storage.Long[newLink.Short] = link.Long
		return storage.Short[newLink.Long], nil
	}

	storage.Long[link.Short] = link.Long
	return storage.Short[link.Long], nil
}

func (storage LocalStorage) GetLongLink(short string) (string, error) {
	if _, ok := storage.Long[short]; ok {
		return storage.Long[short], nil
	}

	return "", fmt.Errorf("для заданной короткой ссылки не существует длинной")
}
