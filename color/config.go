package color

// display style
type style int

const (
	NoStyle      style = iota // 默认
	Bold                      // 加粗
	Faint                     // 弱化
	Italic                    // 斜体
	Underline                 // 下划
	BlinkSlow                 // 慢闪
	BlinkRapid                // 快闪
	ReverseVideo              // 反转
	Concealed                 // 隐藏
	CrossedOut                // 中划
)

// String return description
func (a style) String() string {
	switch a {
	case NoStyle:
		return "默认"
	case Bold:
		return "加粗"
	case Faint:
		return "弱化"
	case Italic:
		return "斜体"
	case Underline:
		return "下划"
	case BlinkSlow:
		return "慢闪"
	case BlinkRapid:
		return "快闪"
	case ReverseVideo:
		return "反转"
	case Concealed:
		return "隐藏"
	case CrossedOut:
		return "中划"
	default:
		return ""
	}
}

// foreground color (Basic 30, Hi-Intensity 90)
type fgColor int

const (
	NoColor fgColor = iota + 89 // 默认
	Black                       // 黑色
	Red                         // 红色
	Green                       // 绿色
	Yellow                      // 黄色
	Blue                        // 蓝色
	Magenta                     // 洋红
	Cyan                        // 青色
	White                       // 白色
)

// String return description
func (fc fgColor) String() string {
	switch fc {
	case NoColor:
		return "默认"
	case Black:
		return "黑色"
	case Red:
		return "红色"
	case Green:
		return "绿色"
	case Yellow:
		return "黄色"
	case Blue:
		return "蓝色"
	case Magenta:
		return "洋红"
	case Cyan:
		return "青色"
	case White:
		return "白色"
	default:
		return ""
	}
}

// background color (Basic 40, Hi-Intensity 100)
type bgColor int

const (
	NoBgColor bgColor = iota + 99 // 默认背景
	BgBlack                       // 黑色背景
	BgRed                         // 红色背景
	BgGreen                       // 绿色背景
	BgYellow                      // 黄色背景
	BgBlue                        // 蓝色背景
	BgMagenta                     // 洋红背景
	BgCyan                        // 青色背景
	BgWhite                       // 白色背景
)

// String return description
func (bc bgColor) String() string {
	switch bc {
	case NoBgColor:
		return "默认背景"
	case BgBlack:
		return "黑色背景"
	case BgRed:
		return "红色背景"
	case BgGreen:
		return "绿色背景"
	case BgYellow:
		return "黄色背景"
	case BgBlue:
		return "蓝色背景"
	case BgMagenta:
		return "洋红背景"
	case BgCyan:
		return "青色背景"
	case BgWhite:
		return "白色背景"
	default:
		return ""
	}
}
