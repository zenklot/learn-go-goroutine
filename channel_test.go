package learn_go_goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Gozenx"
		fmt.Println("Selesai Mengirim")
	}()
	data := <-channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)

}