package color

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestColor(t *testing.T) {
	assert.NotEmpty(t, New("a").Style(Bold).Color(Red).BgColor(BgBlack).Content())

	assert.NotEmpty(t, New("a").Style(1000).Color(1000).BgColor(1000).Content())

	New("a", "b", "c").Color(Green).Render()
}
