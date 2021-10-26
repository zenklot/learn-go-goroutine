package learn_go_goroutine

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello World")
}

func TestCreateGoroutine(t *testing.T) {
	go RunHelloWorld()
	fmt.Println("Ups")
	time.Sleep(1 * time.Second)
}

func DisplayNumber(number int) {
	fmt.Println("Number:", number)
}

func TestNumber(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)

	}
	time.Sleep(10 * time.Second)
}
