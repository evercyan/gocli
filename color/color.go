package color

import (
	"fmt"
	"strings"
)

type color struct {
	text    string
	fgColor fgColor
	bgColor bgColor
	style   style
}

// New return color instance
func New(texts ...string) *color {
	return &color{
		text: strings.Join(texts, " "),
	}
}

// Color set foreground color
func (c *color) Color(v fgColor) *color {
	if v.String() != "" {
		c.fgColor = v
	}
	return c
}

// BgColor set background color
func (c *color) BgColor(v bgColor) *color {
	if v.String() != "" {
		c.bgColor = v
	}
	return c
}

// Style set font style
func (c *color) Style(v style) *color {
	if v.String() != "" {
		c.style = v
	}
	return c
}

// Content return color content
func (c *color) Content() string {
	return fmt.Sprintf("%c[%d;%d;%dm%s%c[0m", 0x1B, c.style, c.bgColor, c.fgColor, c.text, 0x1B)
}

// Render render color content
func (c *color) Render() {
	fmt.Println(c.Content())
}
