package commonresources

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"strings"
)

const (
	maxSmallDescLen = 300
	maxFullDescLen  = 500
	maxSingDescLen  = 0
)

// ---- Get Info -----

// SetLogLevel Sets the log level
func SetLogLevel(log *logrus.Logger,logLevel string) {
	switch strings.ToLower(logLevel) {
	case "trace":
		(*log).SetLevel(logrus.TraceLevel)
		break
	case "debug":
		(*log).SetLevel(logrus.DebugLevel)
		break
	case "info":
		(*log).SetLevel(logrus.InfoLevel)
		break
	case "warn":
		(*log).SetLevel(logrus.WarnLevel)
		break
	case "error":
		(*log).SetLevel(logrus.ErrorLevel)
		break
	}
	(*log).WithField("Level", strings.ToLower(logLevel)).Debug("Log Level Set")
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

//---- Get Download URL ----

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