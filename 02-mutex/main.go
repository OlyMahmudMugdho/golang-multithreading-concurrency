package main

import (
	"fmt"
	"sync"
)

func increment(value *int, mu *sync.Mutex) {
	mu.Lock()
	*value = *value + 1
	defer mu.Unlock()
}

func main() {

	var mu sync.Mutex
	var wg sync.WaitGroup

	var counter int = 0

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			increment(&counter, &mu)
		}()
	}

	wg.Wait()

	fmt.Println(counter)

	/*
		this process will not work perfectly
		var wg sync.WaitGroup

		for i := 0; i < 10000; i++ {
				this process will not work perfectly
				 wg.Add(1)
				go increment(&counter, &wg)
			}

		this process will not work perfectly
		 wg.Wait()
		 fmt.Println(counter)
	*/

}

/*
func increment(value *int, wg *sync.WaitGroup) {
		this process will not work perfectly
		defer wg.Done()
		// *value = *value + 1
	}
*/
