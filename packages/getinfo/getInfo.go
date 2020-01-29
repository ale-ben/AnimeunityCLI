package getinfo

import (
	"github.com/sirupsen/logrus"

	"AnimeunityCLI/packages/commonresources"
)

var (
	//General logger
	log = logrus.New()
	//Log Package wide logger
	Log = log.WithField("Package", "getinfo")
	//File wide logger
	getInfoLog = Log.WithField("File", "getInfo.go")
)

// GetInfo Get a list of anime as a result of a keyword search
func GetInfo(keyword string) []commonresources.AnimeStruct {
	var animeList []commonresources.AnimeStruct

	scraper(keyword, &animeList)
	return animeList
}

// SetLogLevel Sets the log level
func SetLogLevel(logLevel string) {
	commonresources.SetLogLevel(log,logLevel)
}

