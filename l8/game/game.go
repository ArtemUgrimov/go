package game

import (
	"context"
	"encoding/json"
	"fmt"
	"main/l8/quiz"
	"os"
)

type Game struct {
	currentRound int
	Name         string   `json:"GameName"`
	Rounds       []*Round `json:"Rounds"`
	gameOver     bool
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
	go game.CheckForEnd(ctx)
	for {
		nextRound := game.NextRound()
		if nextRound == nil {
			break
		}

		game.PlayRound(nextRound, players)

		if game.gameOver {
			break
		}
	}
	fmt.Println("Game os iver")
	gameOverHandle <- true
}

func (game *Game) PlayRound(round *Round, players map[string]IPlayer) {
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
	<-roundDone
	fmt.Println("Correct answer was", round.RightAnswer, ",-", round.Question.Options[round.RightAnswer])
}

func (g *Game) NextRound() *Round {
	if g.currentRound < len(g.Rounds) {
		g.currentRound++
		return g.Rounds[g.currentRound-1]
	}
	return nil
}

func (g *Game) CheckForEnd(ctx context.Context) {
	<-ctx.Done()
	g.gameOver = true
}
