/**
 * - 不监听系统中断信号, 避免多 goroutine 时, 无法中断整个进程
 * - 在 loading 时, 不对光标进行隐藏处理, 避免意外中断无法对光标进行回显
 * - 偶现重写 line 时会有残余, todo..
 *
 * PS: New 完一定注意要有后续 Success 或 Fail 处理, 避免造成 goroutine 泄漏
 */

package spinner

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/evercyan/gocli/color"
	"github.com/evercyan/gocli/cursor"
)

type spinner struct {
	message  string    // display message
	stopChan chan bool // stop signal
	stopOnce sync.Once // done flag
	loading  []string  // loading style
	symbol   []string  // status symbol
	speed    int       // spinner speed
}

// New return spinner instance
func New(messages ...string) *spinner {
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

// Loading set loading style
func (s *spinner) Loading(loading []string) *spinner {
	if len(loading) > 0 {
		s.loading = loading
	}
	return s
}

// Symbol set status symbol
func (s *spinner) Symbol(symbol []string) *spinner {
	if len(symbol) == 2 {
		s.symbol = symbol
	}
	return s
}

// Loading set loading style
func (s *spinner) Speed(ms int) *spinner {
	if ms < 10 {
		ms = 10
	}
	if ms > 1000 {
		ms = 1000
	}
	s.speed = ms
	return s
}

// Success stop the spinner with success status
func (s *spinner) Success(messages ...string) {
	s.stop(true, messages...)
	return
}

// Fail stop the spinner with fail status
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
			color.New("\r"+s.symbol[0], message).Color(color.Green).Render()
		} else {
			color.New("\r"+s.symbol[1], message).Color(color.Red).Render()
		}
	})
}
