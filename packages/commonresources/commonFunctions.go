package commonresources

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"strings"
)

const (
	maxSmallDescLen = 300
	maxFullDescLen  = 500
)

//TODO: Look for better ways to divide code (Like a placeholder)

// ---- Common ----

//PrintJSONAnimePageStruct Prints an array of anime pages as a JSON
func PrintJSONAnimePageStruct(animePageList []AnimePageStruct) {
	for _, animePage := range animePageList {
		fmt.Printf("{\n \"%s\",\n\"%s\",\n\"%s\",\n[]string{\n", animePage.AnimeID, animePage.AnimeURL, animePage.Title)
		for _, ep := range animePage.EpisodeList {
			fmt.Printf("\"%s\",\n", ep)
		}
		fmt.Printf("},\n%t,\n},", animePage.IsOVA)
	}
}

// ---- Get Info -----

// SetLogLevel Sets the log level for the giver logger
func SetLogLevel(log *logrus.Logger, logLevel string) {
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

/*
PrintAnimeList is used to print a list of animeStruct.

The mode parameter determines the length amount of information to be printed.
This parameter can be:
- 0 for a small description
- 1 for a big description
Any other value results in the max length being set to 0 (Unlimited)
*/
func PrintAnimeList(animeList []AnimeStruct, mode int) {
	for _, anime := range animeList {
		PrintAnime(anime, mode)
	}
}

/*
PrintAnime Prints all the info of an Anime and restrict the description based on the mode parameter.

The mode parameter determines the length amount of information to be printed.
This parameter can be:
- 0 for a small description
- 1 for a big description
Any other value results in the max length being set to 0 (Unlimited)
*/
func PrintAnime(anime AnimeStruct, mode int) {
	var maxDescLen int

	fmt.Printf("\n\n%s - %s\n", anime.AnimeID, anime.Title)

	fmt.Printf("\tEpisodi:  %d\tDurata Episodio: %d\t Durata serie: %d\t Anno uscita: %d\n", anime.NumEpisodes, anime.EpisodeDuration, anime.TotalDuration, anime.Year)

	//Set the max description length based on the mode and print attributes only for determined modes
	if mode == 0 {
		maxDescLen = maxSmallDescLen
	} else {
		fmt.Printf("\tURL\t%s\n\t", "https://animeunity.it/anime.php?id="+anime.AnimeID)
		if mode == 1 {
			maxDescLen = maxFullDescLen
		} else {
			maxDescLen = 0 //Unlimited
		}
	}

	//Reduce the description only if maxDescLen != 0 and the description length is actually greater then the max one
	if maxDescLen != 0 && len(anime.Description) > maxDescLen {
		fmt.Println(anime.Description[:maxDescLen] + "...")
	} else {
		fmt.Println(anime.Description)
	}
}

//---- Get Download URL ----

/*
Unique is used to remove duplicate url from the episode list

Honest to God I have no idea how this works, I found on the web this function to remove duplicates and I used it.
*/
func Unique(animePageList []AnimePageStruct) []AnimePageStruct {
	keys := make(map[string]bool)
	list := []string{}

	for i := 0; i < len(animePageList); i++ { //For every AnimePage in the list passed in input
		for _, entry := range animePageList[i].EpisodeList {
			if _, value := keys[entry]; !value {
				keys[entry] = true
				list = append(list, entry)
			}
		}
		animePageList[i].EpisodeList = list
		list = nil
	}
	return animePageList
}

/*
Sort is used to sort the list of episodes.

Since usually the array is already sorted but it is saved backwards first we invert the array and then we ran it through a bubble sort to check the sorting.
*/
func Sort(animePageList []AnimePageStruct) []AnimePageStruct {
	list := []string{}
	swap := ""

	for i := 0; i < len(animePageList); i++ { //For every AnimePage in the list passed in input

		//Invert the order of the elements on the list
		for j := len(animePageList[i].EpisodeList) - 1; j >= 0; j-- {
			list = append(list, animePageList[i].EpisodeList[j])
		}

		//Bubble sort the list (Should already be in the right order)
		for j := 0; j < len(list)-1; j++ {
			if strings.ToLower(list[j]) > strings.ToLower(list[j+1]) {
				swap = list[j]
				list[j] = list[j+1]
				list[j+1] = swap
				j = 0
			}
		}

		animePageList[i].EpisodeList = list
		list = nil
	}
	return animePageList
}

/*
PrintURLList is used to print the download URLs for an animePage list.

This function calls Unique and Sort on the list before printing it.

The grouping boolean parameter determines if the URLs are going to be splitted in groups or not, if true group size is automatically evaluated based on number of elements and can be 5,10,15,30
*/
func PrintURLList(animePageList []AnimePageStruct, grouping bool) {
	animePageList = Sort(Unique(animePageList))
	for _, animePage := range animePageList {
		i := 0
		groupSize := 0
		arrSize := len(animePage.EpisodeList)

		if arrSize/5 < 10 {
			groupSize = 5
		} else if arrSize/10 < 10 {
			groupSize = 10
		} else if arrSize/15 < 10 {
			groupSize = 15
		} else {
			groupSize = 30
		}

		fmt.Printf("\n\n%s\n", animePage.Title)

		fmt.Printf("Episodes: %d", arrSize)
		if grouping {
			fmt.Printf(", Links per group: %d, Groups: %d\n\n", arrSize/groupSize, groupSize)
		} else {
			fmt.Printf("\n\n")
		}

		for _, url := range animePage.EpisodeList {
			if grouping && i == groupSize {
				fmt.Println("")
				i = 0
			}
			fmt.Println(url)
			i++
		}
	}
	fmt.Println()
}
