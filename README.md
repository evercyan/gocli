<div align="center">

![gocli](https://raw.githubusercontent.com/evercyan/cantor/master/resource/80/8095d74db449976d2ab8020c37afe3e7.png)

self-use go terminal utils

[![travis-ci](https://www.travis-ci.org/evercyan/gocli.svg?branch=master)](https://www.travis-ci.org/evercyan/gocli)
[![codecov](https://codecov.io/gh/evercyan/gocli/branch/master/graph/badge.svg)](https://codecov.io/gh/evercyan/gocli)
[![goreportcard](https://goreportcard.com/badge/github.com/evercyan/gocli)](https://goreportcard.com/report/github.com/evercyan/gocli)

</div>

---

- [Color](##Color) output `color`, `style` text to terminal
- [Table](##Table) output `table` to terminal
- [Spinner](##Spinner)  simple and configurable terminal `spinner`

---

## Color

### Color Usage

```go
package main

import (
    "github.com/evercyan/gocli/color"
)

func main() {
    // font-style:          underline
    // foreground-color:    green
    // background-color:    yellow
    color.New("color").Style(color.Underline).Color(color.Green).BgColor(color.BgYellow).Render()
}
```

[more color example](./examples/color/main.go)

### Color Options

| font-style | foreground-color | background-color |
| --- | --- | --- |
| `color.NoStyle`       | `color.NoColor`   | `color.NoBgColor`     |
| `color.Bold`          | `color.Black`     | `color.BgBlack`       |
| `color.Faint`         | `color.Red`       | `color.BgRed`         |
| `color.Italic`        | `color.Green`     | `color.BgGreen`       |
| `color.Underline`     | `color.Yellow`    | `color.BgYellow`      |
| `color.BlinkSlow`     | `color.Blue`      | `color.BgBlue`        |
| `color.BlinkRapid`    | `color.Magenta`   | `color.BgMagenta`     |
| `color.ReverseVideo`  | `color.Cyan`      | `color.BgCyan`        |
| `color.Concealed`     | `color.White`     | `color.BgWhite`       |
| `color.CrossedOut`    |                   |                       |

### Color Snapshot

![gocli-color](https://raw.githubusercontent.com/evercyan/cantor/master/resource/20/20eaa1795fcee963cb03eaef7fd882fe.jpg)

---

## Table

### Table Usage

```go
package main

import (
    "github.com/evercyan/gocli/table"
)

func main() {
    table.New([][]int{
        {1, 2, 3},
        {4, 5, 6},
        {7, 8, 9},
    }).Style(table.Solid).Border(true).Render()
}
```

[more table example](./examples/table/main.go)

### Table Options

| func | description | eg |
| --- | --- | --- |
| `table.Style(v)`    | set border style  | `table.Solid`<br>`table.Dashed`<br>`table.Dotted` |
| `table.Header(v)`   | set table header  | []string{"col1", "col2"} |
| `table.Border(v)`   | display table data border-bottom  | true |

### Table Snapshot

![gocli-table](https://raw.githubusercontent.com/evercyan/cantor/master/resource/f7/f7c9aaae18aa35f05803c65c4c881267.png)

---

## Spinner

### Spinner Usage

```go
package main

import (
    "time"

    "github.com/evercyan/gocli/spinner"
)

func main() {
    // Loading: set loading style
    // Symbol: set status symbol
    // Speed: set spinner speed(ms)
    s := spinner.New("spinner").Loading(spinner.LoadingA).Symbol(spinner.SymbolA).Speed(100)
    time.Sleep(2 * time.Second)
    s.Success()
}

```

[more spinner example](./examples/spinner/main.go)

### Spinner Options

```go
// loading style
var (
    // LoadingA ...
    LoadingA = []string{"⣾", "⣽", "⣻", "⢿", "⡿", "⣟", "⣯", "⣷"}
    // LoadingB ...
    LoadingB = []string{"\\", "|", "/", "-"}
    // LoadingC ...
    LoadingC = []string{"🌕", "🌖", "🌗", "🌘", "🌑", "🌒", "🌓", "🌔"}
    // LoadingD ...
    LoadingD = []string{"🕐", "🕑", "🕒", "🕓", "🕔", "🕕", "🕖", "🕗", "🕘", "🕙", "🕚", "🕛"}
    // LoadingE ...
    LoadingE = []string{"▁", "▂", "▃", "▄", "▅", "▆", "▇", "█"}
)

// status symbol
var (
    // SymbolA ...
    SymbolA = []string{"✔︎", "✘"}
    // SymbolB ...
    SymbolB = []string{"✓", "✗"}
    // SymbolC ...
    SymbolC = []string{"✅", "❎"}
    // SymbolD ...
    SymbolD = []string{"👍", "👎"}
    // SymbolE ...
    SymbolE = []string{"💚", "💔"}
)
```

### Spinner Snapshot

![gocli-spinner](https://raw.githubusercontent.com/evercyan/cantor/master/resource/d3/d3baabb43de3f6fb925a3f1ef6a92e5a.gif)
