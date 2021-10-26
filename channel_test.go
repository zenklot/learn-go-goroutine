package learn_go_goroutine

import (
	"fmt"
	"strconv"
	"sync"
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

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Gozenx Supriatna"
}
func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	go GiveMeResponse(channel)
	defer close(channel)
	data := <-channel
	fmt.Println(data)
	// time.Sleep(5 * time.Second)

}

func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Raisa Supriatna"
}

func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	go OnlyIn(channel)
	defer close(channel)
	go OnlyOut(channel)
	time.Sleep(5 * time.Second)
}

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	go func() {
		channel <- "Raisa Supriatna"
		channel <- "Gozenx Supriatna"
	}()
	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("Selesai")
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke : " + strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println(data)
	}
	fmt.Println("selesai")
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	var counter int8 = 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data Dari Channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data Dari Channel 2", data)
			counter++
		default:
			fmt.Println("Menunggu Data")
		}
		if counter == 2 {
			break
		}
	}
}

func TestRaceCond(t *testing.T) {
	x := 0

	for i := 1; i < 1000; i++ {
		go func() {
			for j := 1; j < 100; j++ {
				x = x + 1
			}
		}()

	}

	time.Sleep(5 * time.Second)
	fmt.Println("Counter = ", x)
}

func TestMutex(t *testing.T) {
	x := 0
	var mutex sync.Mutex = sync.Mutex{}

	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				mutex.Lock()
				x = x + 1
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Counter = ", x)
}
