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

// DownloadURL Returns a list of episodes download url
func DownloadURL(animePage commonresources.AnimePageStruct, season string) []commonresources.AnimePageStruct {
	downloadURLLog.WithFields(logrus.Fields{
		"animePage": animePage,
		"Season":    season,
	}).Debug("DownloadURL")

	var animePageList []commonresources.AnimePageStruct

	if strings.ToLower(season) != "no" {
		downloadURLLog.WithField("Season", season).Info("Looking for seasons")

		var url string
		if animePage.AnimeID != "" {
			url = "https://animeunity.it/anime.php?id=" + animePage.AnimeID
		} else {
			url = animePage.AnimeURL
		}

		scraper.SeasonScraper(url, strings.ToLower(season), &animePageList)
	} else {
		animePageList = append(animePageList, animePage)
	}

	for i := 0; i < len(animePageList); i++ {
		if animePageList[i].AnimeID != "" {
			animePageList[i].AnimeURL = "https://animeunity.it/anime.php?id=" + animePageList[i].AnimeID
		}
	}

	for i := 0; i < len(animePageList); i++ {
		scraper.EpisodeScraper(&(animePageList[i]))
	}

	return animePageList
}

// SetLogLevel Sets the log level
func SetLogLevel(logLevel string) {
	commonresources.SetLogLevel(log,logLevel)
}
