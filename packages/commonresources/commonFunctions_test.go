package commonresources_test

import (
	"testing"

	"AnimeunityCLI/packages/commonresources"
	"github.com/stretchr/testify/assert"
)

func TestUnique(t *testing.T) {
	assert.Equal(t, correct, commonresources.Unique(original))
}

//TODO: TestSort
func BenchmarkUnique(b *testing.B) {
	// run the DownloadURL function b.N times
	for n := 0; n < b.N; n++ {
		commonresources.SetOwnLogLevel("warn")
		commonresources.Unique(original)
	}
}

func BenchmarkSort(b *testing.B) {
	// run the DownloadURL function b.N times
	for n := 0; n < b.N; n++ {
		commonresources.SetOwnLogLevel("warn")
		commonresources.Sort(original)
	}
}