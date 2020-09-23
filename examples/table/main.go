package main

import (
	"fmt"

	"github.com/evercyan/gocli/color"
	"github.com/evercyan/gocli/table"
)

var (
	structList = []struct {
		Name string `json:"name" table:"header-name"`
		Age  int    `json:"age"`
	}{
		{"Stark", 20},
		{"Lannister", 21},
	}

	numberList = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
)

func main() {
	fmt.Println("-------------------------------- border style")
	for s := table.Solid; s <= table.Dotted; s++ {
		table.New(structList).Style(s).Render()
	}

	fmt.Println("-------------------------------- other")
	fmt.Println("get header from  struct tag `table`:")
	table.New(structList).Render()
	fmt.Println("custom header:")
	table.New(structList).Header([]string{"col1", "col2"}).Border(true).Render()
	fmt.Println("show data border-bottom:")
	table.New(numberList).Border(true).Render()
	fmt.Println("with color:")
	color.New(table.New(structList).Content()).Color(color.Green).Render()
}
