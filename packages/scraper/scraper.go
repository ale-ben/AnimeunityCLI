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

// ---- Get Info ----

//ScrapeInfo query Animeunity with a defined keyword
func ScrapeInfo(keyword string, animeList *[]commonresources.AnimeStruct) {
	scraperLog.WithFields(logrus.Fields{
		"Keyword" : keyword,
		"Anime" : *animeList,
	}).Trace("<ScrapeInfo>")

	url := "https://animeunity.it/anime.php?c=archive"
	// Instantiate default collector
	c := colly.NewCollector(
		// Restrict crawling to specific domains
		colly.AllowedDomains("animeunity.it"), //Prevent the scraper from visiting other websites
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

	// On every a div which has href attribute
	c.OnHTML("div", func(e *colly.HTMLElement) {
		//Check for episode numbers
		if strings.Contains(e.Attr("class"), "card archive-card") { //On every element that has class with card archive-card

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

			anime := commonresources.AnimeStruct{AnimeID: e.ChildAttr("a", "href")[13:], Title: e.ChildText(".card-title"), ImageURL: e.ChildAttr("img", "src"), Description: e.ChildText(".archive-plot"), NumEpisodes: ep, EpisodeDuration: durata, TotalDuration: ep * durata, Year: anno}

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

	scraperLog.WithFields(logrus.Fields{
		"Keyword" : keyword,
		"Anime" : *animeList,
	}).Trace("<ScrapeInfo>")
}

// ---- Get Download URL ----

//SeasonScraper given a Anime Page look for the season list
func SeasonScraper(baseURL string, season string, animePageList *[]commonresources.AnimePageStruct) {
	scraperLog.WithFields(logrus.Fields{
		"Base URL" : baseURL,
		"Season" : season,
	}).Trace("<SeasonScraper>")

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
				requested := false
				seasonInfo := e.ChildText("a[class=text-secondary]")
				year, _ := strconv.Atoi(seasonInfo[0:3])
				isOVA := strings.Contains(strings.ToLower(seasonInfo), "ova") || strings.Contains(strings.ToLower(seasonInfo), "oav")
				scraperLog.WithFields(logrus.Fields{
					"Title": seasonTitle,
					"ID":    seasonID,
					"isOVA": isOVA,
				}).Debug("Season Found")
				if (season == "ova" && isOVA) || (season == "noova" && !isOVA) || (season == "all") || ((season == "no") && (baseURL == "https://animeunity.it/anime.php?id=" + seasonID)){
					requested = true
				}
				if requested {
					animePage := commonresources.AnimePageStruct{AnimeID: seasonID, Title: seasonTitle, EpisodeList: []string{}, IsOVA: isOVA, Year: year}
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

	scraperLog.WithFields(logrus.Fields{
		"Base URL" : baseURL,
		"Season" : season,
	}).Trace("</SeasonScraper>")
}

//EpisodeScraper given an anime page look for the episodes download URL
func EpisodeScraper(animePage *commonresources.AnimePageStruct) {
	scraperLog.WithField("Anime",*animePage).Trace("<EpisodeScraper>")

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
		if strings.Contains(e.Attr("class"), "cus_title") && (*animePage).Title == "" {
			title := (strings.Split(e.Text, " - "))[0]
			scraperLog.WithField("Title", title).Debug("Title Found")
			(*animePage).Title = title
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
		(*animePage).EpisodeList = append((*animePage).EpisodeList, url)
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
	scraperLog.WithField("Anime",*animePage).Trace("</EpisodeScraper>")
}

// SetLogLevel Sets the log level
func SetLogLevel(logLevel string) {
	commonresources.SetLogLevel(log, logLevel,"scraper.go")
}