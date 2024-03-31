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

func (game *Game) GameRunner(ctx context.Context, players map[string]IPlayer, gameOverHandle chan bool) {
	fmt.Printf("Lets play %s\n", game.Name)
	for {
		nextRound := game.NextRound()
		if nextRound == nil {
			fmt.Println("Game over")
			gameOverHandle <- true
			break
		}

		game.PlayRound(ctx, nextRound, players)
	}
}

func (game *Game) PlayRound(ctx context.Context, round *Round, players map[string]IPlayer) {
	game.currentRound++
	fmt.Printf("===> Round %d : <===\n", game.currentRound)

	questions := make(chan *quiz.Question, len(players))
	answersCollector := make(chan *quiz.Answer, len(players))

	for _, p := range players {
		go p.Play(questions, answersCollector)
	}

	go round.Run(questions, len(players))

	var roundDone chan bool = make(chan bool)
	go round.Check(answersCollector, players, roundDone)

	currentCtx, cancel := context.WithTimeout(ctx, time.Duration(time.Second*10))
	defer cancel()
	select {
	case <-roundDone:
		fmt.Println("Correct answer was", round.RightAnswer, ",-", round.Question.Options[round.RightAnswer])
	case <-currentCtx.Done():
		return
	case <-ctx.Done():
		fmt.Println("We want to play but the game is over(")
		return
	}
}

func (g *Game) NextRound() *Round {
	if g.currentRound < len(g.Rounds) {
		g.currentRound++
		return g.Rounds[g.currentRound-1]
	}
	return nil
}
