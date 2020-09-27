package main

import (
	"fmt"

	"github.com/evercyan/gocli/color"
)

func main() {
	fmt.Println("")
	// 显示样式
	for s := color.NoStyle; s <= color.CrossedOut; s++ {
		// 背景颜色
		for bc := color.NoBgColor; bc <= color.BgWhite; bc++ {
			// 字体颜色
			for fc := color.NoColor; fc <= color.White; fc++ {
				fmt.Printf(color.NewColor(
					s.String(),
					bc.String(),
					fc.String(),
				).Style(s).FgColor(fc).BgColor(bc).Content() + " ")
			}
			fmt.Println("")
		}
		fmt.Println("")
	}
	fmt.Println("")
}
