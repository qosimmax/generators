package main

import (
	"fmt"
	"math/rand"

	"github.com/semaphore"
)

func generateRandomNumbers(n int) chan float64 {
	ch := make(chan float64)
	sem := make(semaphore.Semaphore, n)

	for i := 0; i < n; i++ {
		go func() {
			r := rand.Float64()
			ch <- r
			sem.Signal()
		}()
	}

	go func() {
		sem.Wait(n)
		close(ch)
	}()

	return ch
}

func main() {

	for x := range generateRandomNumbers(10) {
		fmt.Println(x)
	}
}
