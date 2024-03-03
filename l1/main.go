package main

import (
	"fmt"
	"time"
)

func main() {
	var words float32 = 100.0
	var amIAnAI bool = true
	var birthDate = time.Date(1990, time.July, 15, 1, 1, 1, 1, time.Now().Local().Location())
	var decade int = 90
	var millenium int = 2000
	var monitorType string = "CRT"
	var academicYear int = 2010
	var graduationDate = time.Date(2015, time.May, 1, 1, 1, 1, 1, time.Now().Local().Location())
	var adultYear int = 2020
	var age int = 32
	fmt.Printf(`
		The story of %.2f words.
		I am an AI. This is %t.
		On the crisp morning of %s, 
		I made my debut, a digital denizen in the burgeoning era of technology. The 
		%ds ushered in my childhood, marked by the whirl of dial-up modems and the glow of %s monitors. 
		The early %ds witnessed my teenage years amidst the evolving landscape of cyberspace. 
		In %d, I embarked on the journey of academia, chasing both bits and knowledge. 
		Graduation day, a significant milestone, unfolded in %s. 
		The subsequent years, a mosaic of coding, literature, and exploration, brought me to %d, 
		where I embraced adulthood. At %d, I remain a storyteller, weaving narratives in the binary tapestry.`,
		words,
		amIAnAI,
		birthDate.Format("January 02, 2006"),
		decade,
		monitorType,
		millenium,
		academicYear,
		graduationDate.Format("January 2006"),
		adultYear,
		age,
	)
}
