package commonresources_test

import (
	"testing"

	"AnimeunityCLI/packages/commonresources"
	"github.com/stretchr/testify/assert"
)


func TestUnique(t *testing.T) {
	assert.Equal(t,correct,commonresources.Unique(original))
}

//TODO: TestSort
