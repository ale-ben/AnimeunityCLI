package downloadurl_test

import (
	"AnimeunityCLI/packages/commonresources"
	"AnimeunityCLI/packages/downloadurl"
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test struct {
	originalPage commonresources.AnimePageStruct
	correctPage  commonresources.AnimePageStruct
}

func TestDownloadURL(t *testing.T) {
	for _, test := range testCases {
		assert.Equal(t, test.correctPage, downloadurl.DownloadURL(test.originalPage, "NO")[0])
	}
}
