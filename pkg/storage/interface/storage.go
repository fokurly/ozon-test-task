package interfacee

import (
	"ozonTest/pkg/models"
)

type Storage interface {
	CreateShortLink(link models.Link) (string, error)
	GetLongLink(short string) (string, error)
}
