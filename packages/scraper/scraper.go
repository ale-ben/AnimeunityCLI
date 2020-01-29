package scraper

import (
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/sirupsen/logrus"

	"AnimeunityCLI/packages/commonresources"
)

var (
	//General logger
	log = logrus.New()
	//Log Package wide logger
	Log = log.WithField("Package", "scraper")
	//File wide logger
	scraperLog = Log.WithField("File", "scraper.go")
)

//TODO: Look for better ways to divide code (Like a placeholder)
//TODO Comment
//TODO GoDoc

// ---- Get Info ----

//ScrapeInfo Query Animeunity with a defined keyword
func ScrapeInfo(keyword string, animeList *[]commonresources.AnimeStruct) {

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

			anime := commonresources.AnimeStruct{e.ChildAttr("a", "href")[13:], e.ChildText(".card-title"), e.ChildAttr("img", "src"), e.ChildText(".archive-plot"), ep, durata, ep * durata, anno}

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

// ---- Get Download URL ----

//SeasonScraper Given a Anime Page look for the season list
func SeasonScraper(baseURL string, season string, animePageList *[]commonresources.AnimePageStruct) {

	// Instantiate default collector
	c := colly.NewCollector(
		// Restrict crawling to specific domains
		colly.AllowedDomains("animeunity.it"),
		// Allow visiting the same page multiple times
		//colly.AllowURLRevisit(),
		// Allow crawling to be done in parallel / async
		//colly.Async(true),
	)

	// Get Seasons
	c.OnHTML("li", func(e *colly.HTMLElement) {
		if strings.Contains(e.Attr("class"), "list-group-item") {
			seasonTitle := e.ChildText("a[href]")
			if len(seasonTitle) != 0 {
				seasonID := e.ChildAttr("a", "href")[13:]
				isOVA := strings.Contains(strings.ToLower(seasonTitle), "ova")
				requested := false
				scraperLog.WithFields(logrus.Fields{
					"Title": seasonTitle,
					"ID":    seasonID,
					"isOVA": isOVA,
				}).Debug("Season Found")
				if (season == "ova" && isOVA) || (season == "noova" && !isOVA) || (season == "all") {
					requested = true
				}
				if requested {
					animePage := commonresources.AnimePageStruct{seasonID, "", seasonTitle, []string{}, isOVA}
					*animePageList = append(*animePageList, animePage)
				}
			}
		}
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		scraperLog.Debug("Visiting ", r.URL.String())
	})

	c.Limit(&colly.LimitRule{
		// Set a delay between requests to these domains
		Delay: 500 * time.Millisecond,
		// Add an additional random delay
		RandomDelay: 1 * time.Second,
	})

	c.Visit(baseURL)
}

//EpisodeScraper Given an anime page look for the episodes download URL
func EpisodeScraper(animePage *commonresources.AnimePageStruct) {

	// Instantiate default collector
	c := colly.NewCollector(
		// Restrict crawling to specific domains
		colly.AllowedDomains("animeunity.it"),
		// Allow visiting the same page multiple times
		//colly.AllowURLRevisit(),
		// Allow crawling to be done in parallel / async
		//colly.Async(true),
	)

	// Get Title
	c.OnHTML("h1", func(e *colly.HTMLElement) {

		//Check for episode numbers
		if strings.Contains(e.Attr("class"), "cus_title") && (*animePage).Titolo == "" {
			title := (strings.Split(e.Text, " - "))[0]
			scraperLog.WithField("Title", title).Debug("Title Found")
			(*animePage).Titolo = title
		}
	})

	// On every a element which has href attribute
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {

		//Check for episode numbers
		if strings.Contains(e.Attr("class"), "ep-button") {
			link := e.Attr("href")
			// Print link
			//fmt.Printf("Link found: %q -> %s\n ", e.Text, link)

			c.Visit(e.Request.AbsoluteURL(link))
		}
	})

	// On every a element which has video attribute
	c.OnHTML("video", func(e *colly.HTMLElement) {

		url := e.ChildAttr("source", "src")
		scraperLog.WithField("URL", url).Debug("Episode Found")
		(*animePage).Episodi = append((*animePage).Episodi, url)
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		scraperLog.Debug("Visiting ", r.URL.String())
	})

	c.Limit(&colly.LimitRule{
		// Set a delay between requests to these domains
		Delay: 500 * time.Millisecond,
		// Add an additional random delay
		RandomDelay: 1 * time.Second,
	})

	c.Visit((*animePage).AnimeURL)
}
