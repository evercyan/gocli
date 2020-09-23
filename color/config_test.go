package color

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestColorConfig(t *testing.T) {
	for s := NoStyle; s <= CrossedOut; s++ {
		assert.NotEmpty(t, s.String())
	}

	for fc := NoColor; fc <= White; fc++ {
		assert.NotEmpty(t, fc.String())
	}

	for bc := NoBgColor; bc <= BgWhite; bc++ {
		assert.NotEmpty(t, bc.String())
	}
}
