package main

import (
	"fmt"
	"math"
	"math/rand"
)

type MinMax struct {
	Min int
	Max int
}

func generator(destination chan int, randeStart int, rangeEnd int, minMax <-chan MinMax, ready chan bool) {
	for i := 0; i < 10; i++ {
		val := randeStart + rand.Intn(rangeEnd-randeStart+1)
		fmt.Println("Generated -> ", val)
		destination <- val
	}
	close(destination)

	mm := <-minMax
	fmt.Println("Min =", mm.Min, "max =", mm.Max)
	ready <- true
}

func minMaxFinder(source <-chan int, destination chan MinMax) {
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

	mm := MinMax{
		Min: min,
		Max: max,
	}

	destination <- mm
	close(destination)
}

func main() {
	randoms := make(chan int)
	minMax := make(chan MinMax)
	wait := make(chan bool)
	go generator(randoms, -10, 10, minMax, wait)
	go minMaxFinder(randoms, minMax)

	<-wait
}
