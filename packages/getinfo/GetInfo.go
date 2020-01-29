package getinfo

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"strings"
)

var (
	maxSmallDescLen = 300
	maxFullDescLen  = 500
	maxSingDescLen  = 0

	//General logger
	log = logrus.New()
	//Log Package wide logger
	Log = log.WithField("Package", "getinfo")
	//File wide logger
	getInfoLog = Log.WithField("File", "GetInfo.go")
)

// AnimeStruct An episode of an anime
type AnimeStruct struct {
	AnimeID        string
	Titolo         string
	ImageURL       string
	Descrizione    string
	Episodi        int
	DurataEpisodio int
	DurataTotale   int
	Anno           int
}

// PrintFullAnimeList Prints all the info of an Anime
func PrintFullAnimeList(animeList []AnimeStruct) {
	for _, anime := range animeList {
		fmt.Printf("\n\n%s - %s\n", anime.AnimeID, anime.Titolo)

		fmt.Printf("\tEpisodi:  %d\tDurata Episodio: %d\t Durata serie: %d\t Anno uscita: %d\n", anime.Episodi, anime.DurataEpisodio, anime.DurataTotale, anime.Anno)

		fmt.Printf("\tURL\t%s\n\t", "https://animeunity.it/anime.php?id="+anime.AnimeID)

		if maxFullDescLen != 0 && len(anime.Descrizione) > maxFullDescLen {
			fmt.Println(anime.Descrizione[:maxFullDescLen] + "...")
		} else {
			fmt.Println(anime.Descrizione)
		}
	}
}

// PrintFullAnime Prints all the info of an Anime
func PrintFullAnime(anime AnimeStruct) {
	fmt.Printf("\n\n%s - %s\n", anime.AnimeID, anime.Titolo)

	fmt.Printf("\tEpisodi:  %d\tDurata Episodio: %d\t Durata serie: %d\t Anno uscita: %d\n", anime.Episodi, anime.DurataEpisodio, anime.DurataTotale, anime.Anno)

	fmt.Printf("\tURL\t%s\n\t", "https://animeunity.it/anime.php?id="+anime.AnimeID)

	if maxSingDescLen != 0 && len(anime.Descrizione) > maxSingDescLen {
		fmt.Println(anime.Descrizione[:maxSingDescLen] + "...")
	} else {
		fmt.Println(anime.Descrizione)
	}
}

// PrintSmallAnimeList Prints a few info on the anime
func PrintSmallAnimeList(animeList []AnimeStruct) {
	for _, anime := range animeList {
		fmt.Printf("\n%s - %s\n\t", anime.AnimeID, anime.Titolo)
		fmt.Printf("Episodi:  %d\tDurata Episodio: %d\t Durata serie: %d\t Anno uscita: %d\n\t", anime.Episodi, anime.DurataEpisodio, anime.DurataTotale, anime.Anno)
		if maxSmallDescLen != 0 && len(anime.Descrizione) > maxSmallDescLen {
			fmt.Println(anime.Descrizione[:maxSmallDescLen] + "...")
		} else {
			fmt.Println(anime.Descrizione)
		}
	}
}

// GetInfo Get a list of anime as a result of a keyword search
func GetInfo(keyword string) []AnimeStruct {

	log.SetLevel(logrus.TraceLevel)
	var animeList []AnimeStruct

	scraper(keyword, &animeList)
	return animeList
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
	getInfoLog.WithField("Level", strings.ToLower(logLevel)).Debug("Log Level Set")
}
