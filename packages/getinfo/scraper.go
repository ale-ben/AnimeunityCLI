package getinfo

import (
	"strconv"
	"strings"
	"time"
	//"net/http"

	"github.com/gocolly/colly/v2"
	//"github.com/gocolly/colly/debug"
)

var (
	//File wide logger
	scraperLog = Log.WithField("File", "scraper.go")
)

func scraper(keyword string, animeList *[]AnimeStruct) {

	url := "https://animeunity.it/anime.php?c=archive"
	// Instantiate default collector
	c := colly.NewCollector(
	// Restrict crawling to specific domains
	//colly.AllowedDomains("animeunity.it"),
	// Allow visiting the same page multiple times
	//colly.AllowURLRevisit(),
	// Allow crawling to be done in parallel / async
	//colly.Async(true),
	//colly.Debugger(&debug.LogDebugger{}),
	)

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		scraperLog.Debug("Visiting ", r.URL.String())
	})

	// On every a element which has href attribute
	c.OnHTML("div", func(e *colly.HTMLElement) {
		//Check for episode numbers
		if strings.Contains(e.Attr("class"), "card archive-card") {

			texts := e.ChildTexts(".card-text")
			var ep, durata, anno int
			for _, text := range texts {
				if strings.Contains(text, "Numero Episodi") {
					buffer := strings.Split(text, ": ")
					ep, _ = strconv.Atoi(buffer[1])
				}

				if strings.Contains(text, "Durata in minuti") {
					buffer := strings.Split(text, ": ")
					durata, _ = strconv.Atoi(buffer[1])
				}

				if strings.Contains(text, "Anno di uscita") {
					buffer := strings.Split(text, ": ")
					anno, _ = strconv.Atoi(buffer[1])
				}

			}

			anime := AnimeStruct{e.ChildAttr("a", "href")[13:], e.ChildText(".card-title"), e.ChildAttr("img", "src"), e.ChildText(".archive-plot"), ep, durata, ep * durata, anno}

			*animeList = append(*animeList, anime)
		}

	})

	c.Limit(&colly.LimitRule{
		// Set a delay between requests to these domains
		Delay: 500 * time.Millisecond,
		// Add an additional random delay
		RandomDelay: 1 * time.Second,
	})

	err := c.Post(url, map[string]string{"query": keyword})

	if err != nil {
		scraperLog.WithField("Error", err).Error("Error in scraper")
	}
}
