package models

type InfoFilm struct {
	Time  string
	Price string
	Class string
}

type FilmModel struct {
	FilmName string
	Genre    string
	Info     []InfoFilm
	RateKP   float64
	RateIMDB float64
	Continue string
	Href     string
	Image    string
}
