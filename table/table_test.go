package table

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type user struct {
	Name string `json:"name" table:"名称"`
	Age  int    `json:"age"`
}

var (
	structList = []user{
		{"Stark", 20},
		{"Lannister", 21},
	}

	structPtrList = []*user{
		{
			"Stark",
			20,
		},
		{
			"Lannister",
			21,
		},
	}

	numberList = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	stringList = [][]string{
		{"a", "bb", "ccc"},
		{"dddd", "eeeee", "ffffff"},
	}
)

func TestTable(t *testing.T) {
	assert.NotEmpty(t, NewTable(structList).Style(Dashed).Border(true).Content())
	assert.NotEmpty(t, NewTable(structList).Style(100).Content())
	assert.NotEmpty(t, NewTable(structList).Header([]string{"Cooooooooooooool1", "Col2"}).Content())
	assert.NotEmpty(t, NewTable(numberList).Header([]string{"Cooooooooooooool1", "Col2", "Col3"}).Content())
	assert.NotEmpty(t, NewTable(stringList).Content())
	NewTable(numberList).Render()
}

func TestTableCoverage(t *testing.T) {
	assert.NotEmpty(t, NewTable(numberList).Style(100).Content())
	assert.NotNil(t, NewTable(1).Content())
	assert.NotNil(t, NewTable(structPtrList).Content())
	assert.NotNil(t, NewTable([]int{1}).Content())
	assert.NotNil(t, NewTable([][]map[string]string{
		{
			{
				"a": "a",
			},
		},
	}).Content())
	assert.NotNil(t, NewTable([][]int{}).Content())
	assert.NotNil(t, NewTable(structList).Header([]string{"c1"}).Content())
	assert.NotNil(t, NewTable(numberList).Header([]string{"c1"}).Content())
}
