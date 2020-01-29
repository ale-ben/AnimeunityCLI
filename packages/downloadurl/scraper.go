package downloadurl

import (
	"strings"
	"time"

	"github.com/gocolly/colly"
	_ "github.com/gocolly/colly/debug"
	"github.com/sirupsen/logrus"
)

var (
	//File wide logger
	scraperLog = Log.WithField("File", "scraper.go")
)

func seScraper(baseURL string, season string, animePageList *([]AnimePageStruct)) {

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
					animePage := AnimePageStruct{seasonID, "", seasonTitle, []string{}, isOVA}
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

func epScraper(animePage *AnimePageStruct) {

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
