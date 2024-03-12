package main

import "fmt"

type Stage struct {
	promptDescription string
	enterDescription  string
	decisions         []*Stage
	effect            func(player *Player)
}

type Player struct {
	name          string
	stage         *Stage
	hp            int
	isBleeding    bool
	extraStrength int
	warmth        int
}

func NewPlayer() *Player {
	player := Player{}
	player.spawn()
	return &player
}

func (player *Player) spawn() {
	fmt.Printf("Як тебе звати, друже? \n")
	fmt.Scanf("%s", &player.name)
	player.stage = nil
	player.hp = 100
	player.isBleeding = false
	player.extraStrength = 0
	player.warmth = 100
}

func (player *Player) updateState() {
	if player.isBleeding {
		player.hp -= 30
	} else {
		player.hp -= 10
	}
	if player.extraStrength >= 50 {
		player.warmth += 20
	}
	player.extraStrength += 10
}

func (player *Player) printState() {
	if player.stage != nil {
		if player.hp < 50 {
			fmt.Printf("%s відчуває різкий біль по всьому тілу, рана від укусу дається взнаки, а в очах періодично темніє\n", player.name)
		}
		if player.isBleeding {
			fmt.Println("Кров постійно тече з рани")
		}
		if player.extraStrength > 30 && player.extraStrength < 70 {
			fmt.Printf("%s відчуває незрозумілий прилив сил, гілки в лісі стали майже пір'ям, а бігати вдається удвічі швидше\n", player.name)
		} else if player.extraStrength >= 70 {
			fmt.Printf("%s може наздогнати будь-яку тварину в лісі, а стрибнути так, що іноді влається спіймати птаха\n", player.name)
		}
		if player.warmth < 50 {
			fmt.Println("{{name}} відчуває сильний холод")
		}
	} else {
		fmt.Printf("%s в лісі, непритомний.\n", player.name)
	}
}

func (stage *Stage) action(player *Player) (res int) {
	if stage.effect != nil {
		stage.effect(player)
	}
	player.stage = stage
	fmt.Println(stage.enterDescription)

	var valuesCount int = 0
	for _, descision := range stage.decisions {
		if descision != nil {
			fmt.Println(descision.promptDescription)
			valuesCount++
		}
	}

	if valuesCount > 0 {
		for res < 1 || res > valuesCount {
			fmt.Printf("Зроби вибір (від 1 до %d)\n", valuesCount)
			fmt.Scan(&res)
			if res < 1 || res > valuesCount {
				player.printState()
			}
		}
	}
	return
}
