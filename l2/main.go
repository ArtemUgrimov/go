package main

import (
	"fmt"
)

func main() {
	monkey1 := Monkey{
		Animal: Animal{
			Name:   "Jija",
			IsFree: false,
		},
		JumpHeight: 25,
	}

	monkey2 := Monkey{
		Animal: Animal{
			Name:   "Chacha",
			IsFree: false,
		},
		JumpHeight: 20,
	}

	l := Lion{
		Animal{
			"Ostap",
			false,
		},
		true,
	}

	giraffe1 := Giraffe{
		Animal: Animal{
			Name:   "Nestor",
			IsFree: false,
		},
		NeckLength: 5,
	}

	giraffe2 := Giraffe{
		Animal: Animal{
			Name:   "Bizon",
			IsFree: false,
		},
		NeckLength: 5,
	}

	cage1 := Cage{&monkey1.Animal}
	cage2 := Cage{&monkey2.Animal}
	cage3 := Cage{&l.Animal}
	cage4 := Cage{&giraffe1.Animal}
	cage5 := Cage{&giraffe2.Animal}
	fmt.Println("All the animals are leaving their cages!!!")

	cage1.freeAnimal()
	cage2.freeAnimal()
	cage3.freeAnimal()
	cage4.freeAnimal()
	cage5.freeAnimal()
	// all the animals left their cages

	fmt.Println(monkey1.whoAreYou())
	fmt.Println(monkey2.whoAreYou())
	fmt.Println(l.whoAreYou())
	fmt.Println(giraffe1.whoAreYou())
	fmt.Println(giraffe2.whoAreYou())

	john := Zookeeper{"John"}
	catchTheAnimal(&john, &monkey1.Animal, &cage1)
	catchTheAnimal(&john, &monkey1.Animal, &cage1) // <= should be an error
	catchTheAnimal(&john, &monkey2.Animal, &cage2)
	catchTheAnimal(&john, &l.Animal, &cage3)
	catchTheAnimal(&john, &l.Animal, &cage2) // <= should be an error
	catchTheAnimal(&john, &giraffe1.Animal, &cage4)
	catchTheAnimal(&john, &giraffe2.Animal, &cage5)
	// zookeeper caught every animal

	fmt.Println(monkey1.whoAreYou())
	fmt.Println(monkey2.whoAreYou())
	fmt.Println(l.whoAreYou())
	fmt.Println(giraffe1.whoAreYou())
	fmt.Println(giraffe2.whoAreYou())
}
