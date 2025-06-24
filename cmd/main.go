package main

import (
	"log"
	"os"
	"strconv"
	"web-scraper/internal/models"
	"web-scraper/internal/processer"
	"web-scraper/internal/services"
	"web-scraper/internal/utils"

	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
)

func setRating(films *[]models.FilmModel, service *services.KinopoiskAPI) error {
	for _, film := range *films {
		KP, IMDB, f_err := service.GetFilmRate(film.FilmName, film.Genre)
		if f_err != nil {
			log.Fatal(f_err)
			return f_err
		}
		film.RateKP, film.RateIMDB = KP, IMDB
	}
	return nil
}

func main() {
	if err := godotenv.Load("./config/.env"); err != nil {
		log.Fatal(err)
	}
	config, c_err := utils.LoadConfig()
	if c_err != nil {
		log.Fatal(c_err)
	}
	cn := cron.New()
	defer cn.Stop()

	_, cn_err := cn.AddFunc(config.CronTime, func() {
		kinopoisk := services.InitKinopoiskInterface(config.KinopoiskURL, os.Getenv("TOKEN"), 10)

		chatID, ch_err := strconv.Atoi(os.Getenv("CHAT_ID"))
		if ch_err != nil {
			log.Fatal(ch_err)
		}
		tg := services.NewTelegramInterface(os.Getenv("BOT_TOKEN"), int64(chatID))
		tomorrowFilms := processer.ScrapingFilm(config.MetropolicTomorrowURL)

		setRating(tomorrowFilms, kinopoisk)

		tg.SendMessage("\n🍿 Расаписание фильмов на завтра в Метрополисе 🍿")
		for _, film := range *tomorrowFilms {
			tg.SendMessage(utils.CreateMessage(film))

		}
	})

	if cn_err != nil {
		log.Println(cn_err)
	}

	cn.Start()
	log.Println("Бот запущен")

	select {}

}
