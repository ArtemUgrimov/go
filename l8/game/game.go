package game

import (
	"context"
	"encoding/json"
	"fmt"
	"main/l8/quiz"
	"os"
	"time"
)

type Game struct {
	currentRound int
	Name         string   `json:"GameName"`
	Rounds       []*Round `json:"Rounds"`
}

func NewGame(rulesFile string) *Game {
	g := &Game{}
	rules, err := os.ReadFile(rulesFile)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	json.Unmarshal(rules, g)
	return g
}

func (game *Game) GameRunner(players map[string]IPlayer, parentCtx context.Context, gameOverHandle context.CancelFunc) {
	fmt.Printf("Lets play %s\n", game.Name)
	for {
		nextRound := game.Next()
		if nextRound == nil {
			fmt.Println("Game over")
			gameOverHandle()
			break
		}
		game.currentRound++
		fmt.Printf("===> Round %d : <===\n", game.currentRound)

		questions := make(chan *quiz.Question, len(players))
		answersCollector := make(chan *quiz.Answer, len(players))

		for _, p := range players {
			go p.Play(questions, answersCollector)
		}

		go nextRound.Run(questions, len(players))

		var done chan bool = make(chan bool)
		go nextRound.Check(answersCollector, players, done)

		ctx, cancel := context.WithTimeout(parentCtx, time.Duration(time.Second*10))
		defer cancel()
		select {
		case <-done:
			fmt.Println("Correct answer was", nextRound.RightAnswer, ",-", nextRound.Question.Options[nextRound.RightAnswer])
		case <-ctx.Done():
			continue
		case <-parentCtx.Done():
			fmt.Println("We want to play but the game is over(")
			return
		}
	}
}

func (g *Game) runCycle(players map[string]IPlayer) {

}

func (g *Game) Next() *Round {
	if g.currentRound < len(g.Rounds) {
		g.currentRound++
		return g.Rounds[g.currentRound-1]
	}
	return nil
}
