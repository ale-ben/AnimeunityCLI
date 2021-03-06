package getinfo

import (
	"github.com/sirupsen/logrus"

	"AnimeunityCLI/packages/commonresources"
	"AnimeunityCLI/packages/scraper"
)

var (
	//General logger
	log = logrus.New()
	//Log Package wide logger
	Log = log.WithField("Package", "getinfo")
	//File wide logger
	getInfoLog = Log.WithField("File", "getInfo.go")
)

//TODO Test Files

// GetInfo get a list of anime as a result of a keyword search
func GetInfo(keyword string) []commonresources.AnimeStruct {
	getInfoLog.WithField("keyword", keyword).Trace("<GetInfo>")
	var animeList []commonresources.AnimeStruct

	scraper.ScrapeInfo(keyword, &animeList)
	getInfoLog.WithField("keyword", keyword).Trace("</GetInfo>")
	return animeList
}

// SetLogLevel Sets the log level
func SetLogLevel(logLevel string) {
	commonresources.SetLogLevel(log, logLevel, "getInfo.go")
}
