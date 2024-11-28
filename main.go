package main

import (
	"fmt"
	"sync"
)

func square(n int, wg *sync.WaitGroup, results chan int) {
	defer wg.Done()
	results <- n * n
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	results := make(chan int, len(numbers))
	var wg sync.WaitGroup

	for _, num := range numbers {
		wg.Add(1)
		go square(num, &wg, results)
	}

	wg.Wait()
	close(results)

	for result := range results {
		fmt.Println(result)
	}
}
