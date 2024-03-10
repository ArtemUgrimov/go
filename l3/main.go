package main

import (
	"fmt"
)

/*
   Можливий сценарій: Стівен прокинувся біля входу в печеру.

*/

func main() {
	//i am using a pointer to not to write & while passing the player as an argument
	var player *Player = &Player{}
	player.spawn()

	start := stage1(player)
	fmt.Println(start.promptDescription)
	playerChoice := start.prompt(player)

	step2 := stage2(player, start, playerChoice)
	playerChoice = step2.prompt(player)

	step3 := stage3(player, step2, playerChoice)
	playerChoice = step3.prompt(player)

	step4 := stage4(player, step3, playerChoice)
	step4.prompt(player)

	finalStage := final(player, step4)
	finalStage.prompt(player)

	player.printState()
}
