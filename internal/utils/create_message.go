package utils

import (
	"fmt"
	"log"
	"strings"
	"web-scraper/internal/models"
)

const Message string = `
👉🏻 Фильм: %s
Жанр: %s
Продолжительность: %s
Рейтинг:
	Кинопоиск: %0.2f
	Imdb: %0.2f
	
%s
`

const MessageInfo string = `
Время сеанса  Ценв билета  Класс зала
%s					%s				%s
`

func CreateMessage(film models.FilmModel) string {
	sb := strings.Builder{}

	_, err := sb.WriteString(fmt.Sprintf(string(Message), film.FilmName, film.Genre, film.Continue, film.RateKP, film.RateIMDB, film.Href))
	if err != nil {
		log.Fatal(err)
	}

	for _, info := range film.Info {
		sb.WriteString(fmt.Sprintf(string(MessageInfo), info.Time, info.Price, info.Class))
	}
	return sb.String()
}
