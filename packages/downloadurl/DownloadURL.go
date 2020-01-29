package downloadurl

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"strings"
)

var (
	//General logger
	log = logrus.New()
	//Log Package wide logger
	Log = log.WithField("Package", "downloadurl")
	//File wide logger
	downloadURLLog = Log.WithField("File", "DownloadURL.go")
)

//AnimePageStruct Struct of an anime page
type AnimePageStruct struct {
	AnimeID  string
	AnimeURL string
	Titolo   string
	Episodi  []string
	IsOVA    bool
}

//Unique Removes duplicates
func Unique(animePageList []AnimePageStruct) []AnimePageStruct {
	keys := make(map[string]bool)
	list := []string{}
	for i := 0; i < len(animePageList); i++ {
		for _, entry := range animePageList[i].Episodi {
			if _, value := keys[entry]; !value {
				keys[entry] = true
				list = append(list, entry)
			}
		}
		animePageList[i].Episodi = list
		list = nil
	}
	return animePageList
}

// PrintURLList Prints a list of URLS
func PrintURLList(animePageList []AnimePageStruct) {
	animePageList = Unique(animePageList)
	for _, animePage := range animePageList {
		i := 0
		groupSize := 0
		arrSize := len(animePage.Episodi)

		if arrSize/5 < 10 {
			groupSize = 5
		} else if arrSize/10 < 10 {
			groupSize = 10
		} else if arrSize/15 < 10 {
			groupSize = 15
		} else {
			groupSize = 30
		}
		fmt.Printf("\n\n%s\n", animePage.Titolo)
		fmt.Printf("Episodes: %d, Links per group: %d, Groups: %d\n\n", arrSize, arrSize/groupSize, groupSize)

		for _, url := range animePage.Episodi {
			if i == groupSize {
				fmt.Println("")
				i = 0
			}
			fmt.Println(url)
			i++
		}
	}
	fmt.Println()
}

// DownloadURL Returns a list of episodes download url
func DownloadURL(animePage AnimePageStruct, season string) []AnimePageStruct {
	downloadURLLog.WithFields(logrus.Fields{
		"animePage": animePage,
		"Season":    season,
	}).Debug("DownloadURL")

	var animePageList []AnimePageStruct

	if strings.ToLower(season) != "no" {
		downloadURLLog.WithField("Season", season).Info("Looking for seasons")

		var url string
		if animePage.AnimeID != "" {
			url = "https://animeunity.it/anime.php?id=" + animePage.AnimeID
		} else {
			url = animePage.AnimeURL
		}

		seScraper(url, strings.ToLower(season), &animePageList)
	} else {
		animePageList = append(animePageList, animePage)
	}

	for i := 0; i < len(animePageList); i++ {
		if animePageList[i].AnimeID != "" {
			animePageList[i].AnimeURL = "https://animeunity.it/anime.php?id=" + animePageList[i].AnimeID
		}
	}

	for i := 0; i < len(animePageList); i++ {
		epScraper(&(animePageList[i]))
	}

	return animePageList
}

// SetLogLevel Sets the log level
func SetLogLevel(logLevel string) {
	switch strings.ToLower(logLevel) {
	case "trace":
		log.SetLevel(logrus.TraceLevel)
		break
	case "debug":
		log.SetLevel(logrus.DebugLevel)
		break
	case "info":
		log.SetLevel(logrus.InfoLevel)
		break
	case "warn":
		log.SetLevel(logrus.WarnLevel)
		break
	case "error":
		log.SetLevel(logrus.ErrorLevel)
		break
	}
	downloadURLLog.WithField("Level", strings.ToLower(logLevel)).Debug("Log Level Set")
}
