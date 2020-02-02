package downloadurl

import (
	"github.com/sirupsen/logrus"
	"strings"

	"AnimeunityCLI/packages/commonresources"
	"AnimeunityCLI/packages/scraper"
)

var (
	//General logger
	log = logrus.New()
	//Log Package wide logger
	Log = log.WithField("Package", "downloadurl")
	//File wide logger
	downloadURLLog = Log.WithField("File", "downloadURL.go")
)

/*
DownloadURL returns a list of episodes download url as a []commonresources.AnimePageStruct.

The season parameter can be:
- OVA (Returns only the URLs for the OVA episodes)
- NOOVA (Returns only the URLs but the OVA episodes)
- ALL (Returns only the URLs for all the seasons)
- NO (Default)(Downloads only the season you pass as URL)
 */
func DownloadURL(animePage commonresources.AnimePageStruct, season string) (animePageList []commonresources.AnimePageStruct) {
	downloadURLLog.WithFields(logrus.Fields{
		"animePage": animePage,
		"Season":    season,
	}).Debug("DownloadURL")

	//Check for seasons if required
	if strings.ToLower(season) != "no" {
		downloadURLLog.WithField("Season", season).Info("Looking for seasons")

		var url string

		//Set the ID or the URL based on the other one
		if animePage.AnimeID != "" {
			url = "https://animeunity.it/anime.php?id=" + animePage.AnimeID
			animePage.AnimeURL = url
		} else {
			url = animePage.AnimeURL
			animePage.AnimeID = url[13:]
		}

		//Update the animePageList with the scraper to get all the seasons and all the info for every season
		scraper.SeasonScraper(url, strings.ToLower(season), &animePageList)

		//If there is only one season and the user required all seasons the scraper would return an empty list because there would be no season section on the website
		if len(animePageList) == 0 {
			animePageList = append(animePageList, animePage)
		}
	} else {
		animePageList = append(animePageList, animePage)
	}

	//For every anime page obtained before update the URL based on the ID provided from the scraper
	for i := 0; i < len(animePageList); i++ {
		if animePageList[i].AnimeID != "" {
			animePageList[i].AnimeURL = "https://animeunity.it/anime.php?id=" + animePageList[i].AnimeID
		}else {
			animePage.AnimeID = animePageList[i].AnimeURL[13:]
		}
	}

	//For each anime page launch the scraper and scrape for episodes URLs
	for i := 0; i < len(animePageList); i++ {
		scraper.EpisodeScraper(&(animePageList[i]))
	}

	return animePageList
}

// SetLogLevel Sets the log level
func SetLogLevel(logLevel string) {
	commonresources.SetLogLevel(log, logLevel)
}
