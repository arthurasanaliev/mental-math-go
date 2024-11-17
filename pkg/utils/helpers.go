package utils

import (
	"fmt"
	"math/rand"
	"time"
)

// ANSI color codes
const (
	Reset = "\033[0m"
	Red   = "\033[31m"
	Green = "\033[32m"
)

// CountDown prints a countdown with a delay
func CountDown() {
	time.Sleep(1 * time.Second)
	for i := 3; i >= 1; i-- {
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}
	fmt.Println("Go!")
	fmt.Println()
	time.Sleep(1 * time.Second)
}

// GetNumbers returns two random numbers in range [l, r]
func GetNumbers(l, r int) (int, int) {
	src := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(src)
	a := rnd.Intn(r-l+1) + l
	b := rnd.Intn(r-l+1) + l
	return a, b
}
