// 终端加载渲染
// 	s := NewSpinner("spinner message")
// 	time.Sleep(time.Second)
// 	s.Success("spinner success")
package spinner

/**
 * - 不监听系统中断信号, 避免多 goroutine 时, 无法中断整个进程
 * - 在 loading 时, 不对光标进行隐藏处理, 避免意外中断无法对光标进行回显
 * - 偶现重写 line 时会有残留多余字符, todo..
 *
 * PS: 初始化完一定要有 Success 或 Fail 处理, 避免造成 goroutine 泄漏
 */

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/evercyan/gocli/color"
	"github.com/evercyan/gocli/cursor"
)

// Spinner interface
type Spinner interface {
	Loading([]string) Spinner // 设置加载样式字符
	Symbol([]string) Spinner  // 设置状态样式字符
	Speed(int) Spinner        // 设置加载速度
	Success(...string)        // 加载成功
	Fail(...string)           // 加载失败
}

type spinner struct {
	message  string    // 显示消息
	stopChan chan bool // 中止信号
	stopOnce sync.Once // 中止标识
	loading  []string  // 加载样式字符
	symbol   []string  // 状态样式字符
	speed    int       // 加载速度
}

// NewSpinner ...
func NewSpinner(messages ...string) Spinner {
	s := &spinner{
		message:  strings.Join(messages, " "),
		stopChan: make(chan bool),
		loading:  LoadingA,
		symbol:   SymbolA,
		speed:    100,
	}
	s.start()
	return s
}

// Loading 设置加载样式字符
func (s *spinner) Loading(loading []string) Spinner {
	if len(loading) > 0 {
		s.loading = loading
	}
	return s
}

// Symbol 设置状态样式字符
func (s *spinner) Symbol(symbol []string) Spinner {
	if len(symbol) == 2 {
		s.symbol = symbol
	}
	return s
}

// Speed 设置加载速度
func (s *spinner) Speed(ms int) Spinner {
	if ms < 10 {
		ms = 10
	}
	if ms > 1000 {
		ms = 1000
	}
	s.speed = ms
	return s
}

// Success 加载成功
func (s *spinner) Success(messages ...string) {
	s.stop(true, messages...)
	return
}

// Fail 加载失败
func (s *spinner) Fail(messages ...string) {
	s.stop(false, messages...)
}

func (s *spinner) start() {
	go func() {
		speed := s.speed
		ticker := time.NewTicker(time.Millisecond * time.Duration(int64(speed)))
		index := 0
		for {
			select {
			case <-ticker.C:
				fmt.Printf("\r%s %s", s.loading[index], s.message)
				index = (index + 1) % len(s.loading)
				if speed != s.speed {
					ticker.Stop()
					speed = s.speed
					ticker = time.NewTicker(time.Millisecond * time.Duration(int64(speed)))
				}
			case <-s.stopChan:
				return
			}
		}
	}()
}

func (s *spinner) stop(status bool, messages ...string) {
	s.stopOnce.Do(func() {
		close(s.stopChan)
		cursor.ClearLine()
		message := s.message
		if len(messages) > 0 {
			message = strings.Join(messages, " ")
		}
		if status {
			color.NewColor("\r"+s.symbol[0], message).FgColor(color.Green).Render()
		} else {
			color.NewColor("\r"+s.symbol[1], message).FgColor(color.Red).Render()
		}
	})
}
