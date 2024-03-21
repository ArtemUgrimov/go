package main

import (
	"fmt"
	"math/rand"
)

func generator(ch chan int) {
	for i := 0; i < 100; i++ {
		ch <- rand.Int() % 100
	}
	close(ch)
}

func avg(input <-chan int, output chan int) {
	sum := 0
	count := 0
	for val := range input {
		sum += val
		count++
	}
	if count > 0 {
		output <- (sum / count)
	} else {
		output <- 0
	}

	close(output)
}

func printer(input <-chan int, ready chan bool) {
	for val := range input {
		fmt.Println("Average is", val)
	}
	ready <- true
	close(ready)
}

func main() {
	randoms := make(chan int)
	average := make(chan int)
	wait := make(chan bool)
	go generator(randoms)
	go avg(randoms, average)
	go printer(average, wait)

	<-wait
}
