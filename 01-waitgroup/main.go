package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now().UnixMilli()

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go HeavyTask(i, &wg)
	}

	wg.Wait()

	end := time.Now().UnixMilli()
	fmt.Printf("taken %v second(s) \n", (end-start)/1000)

}

func HeavyTask(inp int, wg *sync.WaitGroup) {
	time.Sleep(time.Second)
	fmt.Println(inp)
	defer wg.Done()
}
