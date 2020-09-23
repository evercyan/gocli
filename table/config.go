package table

import "errors"

// border style enum
type style int

const (
	Solid  style = iota // 实线
	Dashed              // 虚线(类 mysql 终端表格)
	Dotted              // 点线
)

func (s style) Valid() bool {
	switch s {
	case Solid, Dashed, Dotted:
		return true
	default:
		return false
	}
}

// border struct
type border struct {
	H  rune // 水平 ─
	V  rune // 垂直 │
	VH rune // 水平垂直(相交) ┼
	HU rune // 水平垂直(交上) ┴
	HD rune // 水平垂直(交下) ┬
	VL rune // 垂直水平(交左) ┤
	VR rune // 垂直水平(交右) ├
	DL rune // 转弯(下左) ┐
	DR rune // 转弯(下右) ┌
	UL rune // 转弯(上左) ┘
	UR rune // 转弯(上右) └
}

// border style map
// ref: http://www.tamasoft.co.jp/en/general-info/unicode.html
var styles = map[style]border{
	Solid:  {'─', '│', '┼', '┴', '┬', '┤', '├', '┐', '┌', '┘', '└'},
	Dashed: {'-', '|', '+', '+', '+', '+', '+', '+', '+', '+', '+'},
	Dotted: {'*', '*', '*', '*', '*', '*', '*', '*', '*', '*', '*'},
}

var chineseCharset = []rune{0x2E80, 0x9FD0}

var (
	errHeader = errors.New("invalid header length")
	errType   = errors.New("only support slice or array")
)