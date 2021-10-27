package learn_go_goroutine

import (
	"fmt"
	"sync"
	"testing"
)

func AddToMap(data *sync.Map, value int, group *sync.WaitGroup) {
	data.Store(value, value)
	group.Done()
}

func TestSyncMap(t *testing.T) {
	var data *sync.Map = &sync.Map{}
	var group *sync.WaitGroup = &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		group.Add(1)
		go AddToMap(data, i, group)
	}

	group.Wait()

	data.Range(func(key, value interface{}) bool {
		fmt.Println(key, ":", value)
		return true
	})

	fmt.Println("selesai")
}
