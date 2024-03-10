package main

func countTwoCharacters(input string, char1, char2 rune) (firstCount int32, secondCount int32) {
	for _, val := range input {
		switch val {
		case char1:
			firstCount++
		case char2:
			secondCount++
		}
	}
	return
}

func main() {
	/*
	   Можливий сценарій: Стівен прокинувся біля входу в печеру.
	   Він лише памʼятає своє імʼя. Поряд з ним рюкзак, в якому він знаходить сірники, ліхтарик і ніж.
	   У печері темно, тому Стівен іде стежкою, яка веде від печери в ліс.
	   У лісі Стівен натикається на мертве тіло дивної тварини. Він обирає нічого з цим не робити й іти далі.
	   Через деякий час Стівен приходить до безлюдного табору.
	   Він вже втомлений і вирішує відпочити, а не йти далі.
	   У найближчому наметі він знаходить сейф з кодовим замком з двох чисел.
	   Він добирає код, і коли сейф відчиняється, йому на долоню виповзає велика комаха, кусає його й тікає.
	   Стівен непритомніє. А все могло бути зовсім інакше.
	*/

	// char1, char2 := 'a', 'c'
	// text := "abcd abc cd"
	// count1, count2 := countTwoCharacters(text, char1, char2)
	// fmt.Printf("character `%c` occured %d times, character `%c` occured %d times\n", char1, count1, char2, count2)

	// N := 30
	// K := 10
	// for i := 1; i < N; i++ {
	// 	if i%K == 0 {
	// 		fmt.Println(i)
	// 	}
	// }

	// a := []int{10, 20, 333, 0, -50}
	// for idx, elem := range a {
	// 	fmt.Println(idx, elem)
	// }

	// for _, elem := range a {
	// fmt.Println(elem)
	// }

	/*
		name := "Artem"
		nameLen := len(name)

		switch nameLen {
		case 5:
			fmt.Println("Artem's name length is 5")
		case 6:
			fmt.Println("???")
		default:
			fmt.Println("Something went wrong")
		}

		for i := 0; i < 100; i++ {
			rnd := rand.Int() % 10
			if rnd == 5 {
				fmt.Printf("Rnd = %d. Break\n", rnd)
				break
			} else if rnd == 7 {
				continue
			}

			fmt.Printf("%dth iteation\n", i)
		}

		for pos, char
	*/
}
