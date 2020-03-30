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
	}).Trace("<DownloadURL>")

	downloadURLLog.WithFields(
		logrus.Fields{
			"Season": season,
			"Anime":  animePage,
		}).Info("Looking for seasons")

	var url string

	//Set the ID or the URL based on the other one
	if animePage.AnimeID != "" {
		url = "https://animeunity.it/anime.php?id=" + animePage.AnimeID
		animePage.AnimeURL = url
	} else {
		url = animePage.AnimeURL
		animePage.AnimeID = url[35:]
	}

	downloadURLLog.WithField("Anime", animePage).Debug("Updated basic Season Info")

	//Update the animePageList with the scraper to get all the seasons and all the info for every season
	scraper.SeasonScraper(url, strings.ToLower(season), &animePageList)

	if log.GetLevel() == logrus.DebugLevel || log.GetLevel() == logrus.TraceLevel {
		for _, animeP := range animePageList {
			downloadURLLog.WithFields(
				logrus.Fields{
					"Season":      season,
					"Anime Pages": animeP,
				}).Info("Season Scraping Completed")
		}
	}

	//If there is only one season and the user required all seasons the scraper would return an empty list because there would be no season section on the website
	if len(animePageList) == 0 {
		downloadURLLog.WithField("Season", season).Debug("No other season found, using default URL")
		animePageList = append(animePageList, animePage)
	}

	//For every anime page obtained before update the URL based on the ID provided from the scraper
	for i := 0; i < len(animePageList); i++ {
		if animePageList[i].AnimeID != "" {
			animePageList[i].AnimeURL = "https://animeunity.it/anime.php?id=" + animePageList[i].AnimeID
		} else {
			animePage.AnimeID = animePageList[i].AnimeURL[13:]
		}
		downloadURLLog.WithField("Updated Anime", animePageList[i]).Debug("Updating Anime Info")
	}

	downloadURLLog.WithField("Ova",animePageList[0].IsOVA).Error("OVA")

	//For each anime page launch the scraper and scrape for episodes URLs
	for i := 0; i < len(animePageList); i++ {
		downloadParallAux(&animePageList[i]) //TODO: Fix
	}

	downloadURLLog.WithFields(logrus.Fields{
		"animePage": animePage,
		"Season":    season,
	}).Trace("</DownloadURL>")
	return animePageList
}

func downloadParallAux(animePage *commonresources.AnimePageStruct) {
	downloadURLLog.WithField("Anime", *animePage).Trace("<downloadParallAux>")
	scraper.EpisodeScraper(animePage)
	downloadURLLog.WithField("Anime", *animePage).Info("Episode Scraping Completed")
	downloadURLLog.WithField("Anime", *animePage).Trace("</downloadParallAux>")
}

// SetLogLevel Sets the log level
func SetLogLevel(logLevel string) {
	commonresources.SetLogLevel(log, logLevel, "downloadURL.go")
}
