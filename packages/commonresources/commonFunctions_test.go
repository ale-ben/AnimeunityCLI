package commonresources_test

import (
	"testing"

	"AnimeunityCLI/packages/commonresources"
	"github.com/stretchr/testify/assert"
)

/*
//AnimePageStruct Struct of an anime page
type AnimePageStruct struct {
	AnimeID  string
	AnimeURL string
	Titolo   string
	Episodi  []string
	IsOVA    bool
}
*/

func TestUnique(t *testing.T) {
	assert.Equal(t,correct,commonresources.Unique(original))
}

/*

 */
