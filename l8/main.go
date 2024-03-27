package main

import (
	"context"
	"fmt"
	"main/l8/game"
	"main/l8/player"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	playersCount := 5
	players := make(map[string]game.IPlayer)
	for i := 0; i < playersCount; i++ {
		p := player.NewPlayer()
		players[p.Name] = p
	}

	mainCtx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(time.Second*100))
	defer cancelFunc()

	game := game.NewGame("game.json")
	go game.GameRunner(players, mainCtx, cancelFunc)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-mainCtx.Done():
	case <-sigs:
		cancelFunc()
	}

	winner := getWinner(players)
	fmt.Printf("👑👑👑The winner is %s with the score of %d! Congratulations!👑👑👑\n\n", winner.GetName(), winner.GetScore())

	fmt.Println("Playground is closed")
}

func getWinner(players map[string]game.IPlayer) game.IPlayer {
	var winner game.IPlayer = nil
	for _, p := range players {
		if winner == nil || p.GetScore() > winner.GetScore() {
			winner = p
		}
	}
	return winner
}
