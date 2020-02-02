//	Package commonresources provides common structs and function to the whole program
package commonresources

/*
AnimeStruct contains the overall info on an anime season.

This is used when searching an anime.
 */
type AnimeStruct struct {
	AnimeID         string
	Title           string
	ImageURL        string
	Description     string
	NumEpisodes     int //Number of episodes in the season
	EpisodeDuration int //Duration of a single episode in minutes
	TotalDuration   int //Overall duration in minutes of the season (EpisodeDuration * NumEpisodes)
	Year            int
}

/*
AnimePageStruct contains all the specific info of an anime season, including download links.

This is used when looking for anime episodes.
*/
type AnimePageStruct struct {
	AnimeID     string
	AnimeURL    string
	Title       string
	EpisodeList []string
	IsOVA       bool
}
