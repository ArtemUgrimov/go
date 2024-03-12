package main

import (
	"fmt"
)

/*
   Можливий сценарій: Стівен прокинувся біля входу в печеру.

*/

func main() {
	//i am using a pointer to not to write & while passing the player as an argument
	player := NewPlayer()

	start := stage1(player)
	fmt.Println(start.promptDescription)
	playerChoice := start.action(player)

	step2 := stage2(player, start, playerChoice)
	playerChoice = step2.action(player)

	step3 := stage3(player, step2, playerChoice)
	playerChoice = step3.action(player)

	step4 := stage4(player, step3, playerChoice)
	step4.action(player)

	finalStage := final(player, step4)
	finalStage.action(player)

	player.printState()
}
