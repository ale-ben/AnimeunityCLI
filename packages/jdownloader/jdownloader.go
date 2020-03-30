package jdownloader

import (
	"AnimeunityCLI/packages/commonresources"
	"github.com/sirupsen/logrus"
	"path"
	"strings"

	"path/filepath"
)

var (
	//General logger
	log = logrus.New()
	//Log Package wide logger
	Log = log.WithField("Package", "jdownloader")
	//File wide logger
	jdownloaderLog = Log.WithField("File", "getInfo.go")
)

//TODO Test Files
//TODO Comment

func formatTitle(oldTitile string) (newTitle string){

	//Formatting the string
	newTitle = strings.ReplaceAll(oldTitile, " ", "_")
	newTitle = strings.ReplaceAll(newTitle, "!", "")
	newTitle = strings.ReplaceAll(newTitle, ":", "")
	newTitle = strings.ReplaceAll(newTitle, ",", "")

	return
}

func createCrawlFile(animePage commonresources.AnimePageStruct, crawlPath string, jdownloadPath string) (err error) {
	jdownloaderLog.WithFields(logrus.Fields{
		"Anime": animePage,
		"Crawl Path": crawlPath,
		"Download Path" : jdownloadPath,
	}).Trace("<createCrawlPath>")

	fileContent := ""

	//Formatting the string
	formattedAnimeTitle := formatTitle(animePage.Title)

	//Creating the AnimeDir
	animeDir := filepath.Join(jdownloadPath, formattedAnimeTitle)

	for _, ep := range animePage.EpisodeList {
		fileContent += "{\n"
		fileContent += "\ttext= " + ep + "\n"
		fileContent += "\tdownloadFolder= " + animeDir + "\n"
		fileContent += "\tenabled= true\n"
		fileContent += "\tautoStart= true\n"
		fileContent += "\tautoConfirm= true\n"
		fileContent += "}\n"
	}
	jdownloaderLog.WithFields(logrus.Fields{
		"File Path" : crawlPath,
		"File Title" : formattedAnimeTitle+".crawljob",
		"File Content" : fileContent,
	}).Debug("Writing to file")
	err = commonresources.WriteToFile(crawlPath, formattedAnimeTitle+".crawljob", fileContent)

	jdownloaderLog.WithFields(logrus.Fields{
		"Anime": animePage,
		"Crawl Path": crawlPath,
		"Download Path" : jdownloadPath,
	}).Trace("</createCrawlPath>")
	return
}

func SendToJDownloader(animePageList []commonresources.AnimePageStruct, crawlPath string, jdownloadPath string) (err error) {
	jdownloaderLog.WithFields(logrus.Fields{
		"Crawl Path": crawlPath,
		"Download Path" : jdownloadPath,
	}).Trace("<SendToJDownloader>")

	if len(animePageList)>1 {
		formattedTitle := formatTitle(animePageList[0].Title)
		jdownloadPath = path.Join(jdownloadPath,formattedTitle)
		jdownloaderLog.WithField("New Path",jdownloadPath).Debug("Multiple Seasons detected, updating download path")
	}
	for _, animePage := range animePageList {
		jdownloaderLog.WithField("Anime",animePage).Debug("Creating file for Anime")
		err = createCrawlFile(animePage, crawlPath, jdownloadPath)
		if err != nil {
			jdownloaderLog.WithField("Error", err).Warn("Error while creating crawljobs")
		}
	}

	jdownloaderLog.WithFields(logrus.Fields{
		"Crawl Path": crawlPath,
		"Download Path" : jdownloadPath,
	}).Trace("</SendToJDownloader>")
	return
}

// SetLogLevel Sets the log level
func SetLogLevel(logLevel string) {
	commonresources.SetLogLevel(log, logLevel)
}
