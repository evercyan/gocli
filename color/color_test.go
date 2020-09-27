package color

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestColor(t *testing.T) {
	assert.NotEmpty(t, NewColor("a").Style(Bold).FgColor(Red).BgColor(BgBlack).Content())

	assert.NotEmpty(t, NewColor("a").Style(1000).FgColor(1000).BgColor(1000).Content())

	NewColor("a", "b", "c").FgColor(Green).Render()
}
