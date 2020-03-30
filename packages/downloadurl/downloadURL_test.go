package downloadurl_test

import (
	"AnimeunityCLI/packages/commonresources"
	"AnimeunityCLI/packages/downloadurl"
	"github.com/stretchr/testify/assert"
	"testing"
)

var(
	anime = commonresources.AnimePageStruct{
		AnimeID:     "200",
		AnimeURL:    "https://animeunity.it/anime.php?id=200",
		Title:       "",
		EpisodeList: nil,
		IsOVA:       false,
		Year:        0,
	}
)

type Test struct {
	originalPage commonresources.AnimePageStruct
	correctPage  commonresources.AnimePageStruct
}

func TestDownloadURL(t *testing.T) {
	for _, test := range testCases {
		downloadurl.SetLogLevel("warn")
		assert.Equal(t, test.correctPage, downloadurl.DownloadURL(test.originalPage, "NO")[0])
	}
}

func benchmarkDownloadURL(anime commonresources.AnimePageStruct,season string, b *testing.B) {
	// run the DownloadURL function b.N times
	for n := 0; n < b.N; n++ {
		downloadurl.SetLogLevel("warn")
		downloadurl.DownloadURL(anime,season)
	}
}

func BenchmarkDownloadURLALL(b *testing.B) {
	benchmarkDownloadURL(anime,"all",b)
}

func BenchmarkDownloadURLNO(b *testing.B) {
	benchmarkDownloadURL(anime,"no",b)
}

func BenchmarkDownloadURLOVA(b *testing.B) {
	benchmarkDownloadURL(anime,"ova",b)
}

func BenchmarkDownloadURLNOOVA(b *testing.B) {
	benchmarkDownloadURL(anime,"noova",b)
}