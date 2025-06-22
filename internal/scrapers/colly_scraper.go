package scrapers

import (
	"github.com/gocolly/colly"
)

func InitCollyScrapper() *colly.Collector {
	scraper := *colly.NewCollector()
	scraper.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36"
	return &scraper
}
