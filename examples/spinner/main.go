package main

import (
	"fmt"
	"time"

	"github.com/evercyan/gocli/spinner"
)

func basic() {
	fmt.Println("\n-------------------------------- basic")

	s1 := spinner.New("success message")
	time.Sleep(1 * time.Second)
	s1.Success("success")

	s2 := spinner.New("fail message")
	time.Sleep(1 * time.Second)
	s2.Fail("fail")
}

func speed() {
	fmt.Println("\n-------------------------------- speed")

	s := spinner.New("reset speed")
	time.Sleep(3 * time.Second)
	s.Speed(10)
	time.Sleep(3 * time.Second)
	s.Success()
}

func loading() {
	fmt.Println("\n-------------------------------- loading")

	loadings := [][]string{
		spinner.LoadingA,
		spinner.LoadingB,
		spinner.LoadingC,
		spinner.LoadingD,
		spinner.LoadingE,
	}
	for _, loading := range loadings {
		sl := spinner.New(loading...).Loading(loading)
		time.Sleep(2 * time.Second)
		sl.Success()
	}
}

func symbol() {
	fmt.Println("\n-------------------------------- symbol")

	symbols := [][]string{
		spinner.SymbolA,
		spinner.SymbolB,
		spinner.SymbolC,
		spinner.SymbolD,
		spinner.SymbolE,
	}
	for _, symbol := range symbols {
		s1 := spinner.New(symbol...).Symbol(symbol)
		s2 := spinner.New(symbol...).Symbol(symbol)
		time.Sleep(2 * time.Second)
		s1.Success()
		s2.Fail()
	}
}

func main() {
	basic()
	speed()
	loading()
	symbol()
}
