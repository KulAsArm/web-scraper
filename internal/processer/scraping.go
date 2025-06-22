package processer

import (
	"strings"
	"web-scraper/internal/models"
	"web-scraper/internal/scrapers"

	"github.com/gocolly/colly"
)

func ScrapingFilm(url string) *[]models.FilmModel {
	scrap := scrapers.InitCollyScrapper()
	films := []models.FilmModel{}

	scrap.OnHTML(".shedule_content .shedule_movie", func(e *colly.HTMLElement) {
		genre := e.Attr("data-gtm-list-item-genre")

		href := e.ChildAttr(".gtm-ec-list-item-movie", "href")
		name := e.ChildAttr(".gtm-ec-list-item-movie", "data-gtm-ec-name")

		image := e.ChildAttr(".shedule_movie_content img", "data-src")

		cont := strings.Split(strings.ReplaceAll(e.ChildText(".shedule_movie_description .title"), " ", ""), "\n")[2]

		times := []string{}
		prices := []string{}
		classes := []string{}
		e.ForEach(".shedule_movie_sessions .buy_seance", func(i int, h *colly.HTMLElement) {
			times = append(times, h.ChildText(".shedule_session_time"))
			prices = append(prices, h.ChildText(".shedule_session_price"))
			classes = append(classes, strings.ReplaceAll(h.ChildText(".shedule_session_format"), "     ", ""))

		})

		info := []models.InfoFilm{}
		for index, value := range times {
			info = append(info, models.InfoFilm{
				Time:  value,
				Price: prices[index],
				Class: classes[index],
			})
		}
		kp, imdb := 0.0, 0.0

		films = append(films, models.FilmModel{
			FilmName: name,
			Genre:    genre,
			Info:     info,
			RateKP:   kp,
			RateIMDB: imdb,
			Continue: cont,
			Href:     href,
			Image:    image,
		})

	})
	scrap.Visit(url)
	return &films
}
