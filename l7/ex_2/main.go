package main

import (
	"fmt"
	"math"
	"math/rand"
)

func generator(destination chan int, randeStart int, rangeEnd int, minMax <-chan int, ready chan bool) {
	for i := 0; i < 10; i++ {
		val := randeStart + rand.Intn(rangeEnd-randeStart+1)
		fmt.Println("Generated -> ", val)
		destination <- val
	}
	close(destination)

	min := <-minMax
	max := <-minMax
	fmt.Println("Min =", min, "max =", max)
	ready <- true
	close(ready)
}

func minMaxFinder(source <-chan int, destination chan int) {
	min := math.MaxInt32
	max := math.MinInt32
	for num := range source {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}
	destination <- min
	destination <- max
	close(destination)
}

func main() {
	randoms := make(chan int)
	minMax := make(chan int)
	wait := make(chan bool)
	go generator(randoms, -10, 10, minMax, wait)
	go minMaxFinder(randoms, minMax)

	<-wait
}
