package jdownloader

import (
	"AnimeunityCLI/packages/commonresources"
	"github.com/sirupsen/logrus"
	"strings"

	"path/filepath"
)

var (
	//General logger
	log = logrus.New()
	//Log Package wide logger
	Log = log.WithField("Package", "getinfo")
	//File wide logger
	jdownloaderLog = Log.WithField("File", "getInfo.go")
)

//TODO Test Files
//TODO Comment

func createCrawlFile(animePage commonresources.AnimePageStruct, crawlPath string, jdownloadPath string) (err error) {
	fileContent := ""

	//Formatting the string
	formattedAnimeTitle := strings.ReplaceAll(animePage.Title," ","_")
	formattedAnimeTitle = strings.ReplaceAll(formattedAnimeTitle,"!","")
	formattedAnimeTitle = strings.ReplaceAll(formattedAnimeTitle,":","")
	formattedAnimeTitle = strings.ReplaceAll(formattedAnimeTitle,",","")

	//Creating the AnimeDir
	animeDir := filepath.Join(jdownloadPath,formattedAnimeTitle)

	for _, ep := range animePage.EpisodeList {
		fileContent += "{\n"
		fileContent += "\ttext= " + ep + "\n"
		fileContent += "\tdownloadFolder= " + animeDir + "\n"
		fileContent += "\tenabled= true\n"
		fileContent += "\tautoStart= true\n"
		fileContent += "\tautoConfirm= true\n"
		fileContent += "}\n"
	}

	err = commonresources.WriteToFile(crawlPath,formattedAnimeTitle + ".crawljob",fileContent)

	return
}

func SendToJDownloader(animePageList []commonresources.AnimePageStruct, crawlPath string, jdownloadPath string) (err error){
	for _, animePage := range animePageList {
		err = createCrawlFile(animePage,crawlPath,jdownloadPath)
		if err != nil {
			jdownloaderLog.WithField("Error",err).Warn("Error while creating crawljobs")
		}
	}

	return
}

// SetLogLevel Sets the log level
func SetLogLevel(logLevel string) {
	commonresources.SetLogLevel(log, logLevel)
}