package spinner

import (
	"testing"
	"time"
)

func TestSpinner(t *testing.T) {
	s1 := New("spinner1 message").Loading(LoadingE).Symbol(SymbolE).Speed(200)
	s1.Speed(1)
	time.Sleep(time.Millisecond * 10)
	s1.Speed(10000)
	time.Sleep(time.Millisecond * 10)
	s1.Success("spinner1 success")

	s2 := New("spinner2 message")
	time.Sleep(time.Millisecond * 10)
	s2.Fail("spinner2 fail")
	s2.Fail("spinner2 fail2")
}
