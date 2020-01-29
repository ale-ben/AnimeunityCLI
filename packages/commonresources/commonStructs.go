package commonresources

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

//AnimePageStruct Struct of an anime page
type AnimePageStruct struct {
	AnimeID  string
	AnimeURL string
	Titolo   string
	Episodi  []string
	IsOVA    bool
}