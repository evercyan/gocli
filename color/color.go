// 终端文本前景色色, 背景色, 显示样式渲染
// 	color.NewColor("text").FgColor(color.Green).Render()
package color

import (
	"fmt"
	"strings"
)

// Color interface
type Color interface {
	FgColor(FgColor) Color // 设置前景色
	BgColor(BgColor) Color // 设置背景色
	Style(Style) Color     // 设置样式
	Content() string       // 返回渲染后文本内容
	Render()               // 输出渲染后文本内容
}

type color struct {
	text    string
	fgColor FgColor
	bgColor BgColor
	style   Style
}

// NewColor ...
func NewColor(texts ...string) Color {
	return &color{
		text: strings.Join(texts, " "),
	}
}

// FgColor 设置前景色
func (c *color) FgColor(v FgColor) Color {
	if v.String() != "" {
		c.fgColor = v
	}
	return c
}

// BgColor 设置背景色
func (c *color) BgColor(v BgColor) Color {
	if v.String() != "" {
		c.bgColor = v
	}
	return c
}

// Style 设置样式
func (c *color) Style(v Style) Color {
	if v.String() != "" {
		c.style = v
	}
	return c
}

// Content 返回渲染后文本内容
func (c *color) Content() string {
	return fmt.Sprintf("%c[%d;%d;%dm%s%c[0m", 0x1B, c.style, c.bgColor, c.fgColor, c.text, 0x1B)
}

// Render 输出渲染后文本内容
func (c *color) Render() {
	fmt.Println(c.Content())
}
