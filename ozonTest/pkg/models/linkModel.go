package models

import (
	"math/rand"
	"time"
)

type Link struct {
	Short string `json:"shortLink"`
	Long  string `json:"longLink"`
}

const symbols = "_123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func (l *Link) turnToShort() {
	rand.Seed(time.Now().UnixNano())
	link := make([]byte, 10)
	for i := range link {
		link[i] = symbols[rand.Intn(len(symbols))]
	}

	l.Short = string(link)
}

func NewLink(url string) Link {
	myLink := Link{Long: url}
	myLink.turnToShort()

	return myLink
}
