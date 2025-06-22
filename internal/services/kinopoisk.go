package services

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type KinopoiskAPI struct {
	Url   string
	Limit int
	Token string
}

type rating struct {
	KP   float64 `json:"kp"`
	Imdb float64 `json:"imdb"`
}

type movie struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Type   string `json:"type"`
	Year   string `json:"year"`
	Rating rating `json:"rating"`
}

type response struct {
	Data []movie `json:"data"`
}

func (k *KinopoiskAPI) GetFilmRate(name, _type string) (float64, float64, error) {
	params := url.Values{}
	params.Set("page", "1")
	params.Set("limit", strconv.Itoa(k.Limit))
	params.Set("query", name)

	searchUrl := k.Url + params.Encode()

	req, req_err := http.NewRequest(http.MethodGet, searchUrl, nil)
	if req_err != nil {
		log.Fatal(req_err)
	}
	req.Header.Add("X-API-KEY", k.Token)
	req.Header.Add("accept", "application/json")

	client := &http.Client{Timeout: 5 * time.Second}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	response := new(response)

	err = json.NewDecoder(resp.Body).Decode(&response)

	kp := 0.0
	imdb := 0.0

	defer resp.Body.Close()
	for _, value := range response.Data {
		if value.Name == name && value.Type == _type {
			return value.Rating.KP, value.Rating.Imdb, nil
		}
	}
	return kp, imdb, err
}

func InitKinopoiskInterface(url, token string, limit int) *KinopoiskAPI {
	service := KinopoiskAPI{
		Url:   url,
		Limit: limit,
		Token: token,
	}
	return &service
}
