package main

import (
	"fmt"
	"sync"
)

func numberTransformer(x int) int  {
	return x * x
}

func reader(id int, channel chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		val, ok := <-channel
		if !ok {
			return
		}
		fmt.Printf("Reader %d is processing channel %d\n", id, numberTransformer(val))
	}
}

func main() {
	var wg sync.WaitGroup
	channel := make(chan int)

	wg.Add(4)

	go reader(1, channel, &wg)
	go reader(2, channel, &wg)
	go reader(3, channel, &wg)
	go reader(4, channel, &wg)

	for i := 0; i < 100; i++ {
		channel <- i
	}

	close(channel)

	wg.Wait()

}