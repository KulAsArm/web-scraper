package utils

import (
	"log"
	"os"

	yaml "gopkg.in/yaml.v3"
)

type config struct {
	MetropolicTodayURL    string `yaml:"metropolic_today"`
	MetropolicTomorrowURL string `yaml:"metropolic_tomorrow"`
	KinopoiskURL          string `yaml:"kinopoisk_url"`
	CronTime              string `yaml:"cron_time"`
}

func LoadConfig() (*config, error) {

	config := new(config)

	yamlFile, err := os.ReadFile("./config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatal(err)
	}
	return config, err
}
