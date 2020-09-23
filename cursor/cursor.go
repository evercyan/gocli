package cursor

import (
	"fmt"
)

// HideCursor ...
func HideCursor() {
	fmt.Printf("\033[?25l")
}

// ShowCursor ...
func ShowCursor() {
	fmt.Printf("\033[?25h")
}

// ClearLine ...
func ClearLine() {
	fmt.Printf("\r\033[0K")
}
