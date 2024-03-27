package player

import (
	"fmt"
	"html"
	"main/l8/quiz"
	"math/rand"
	"strconv"
	"time"
)

type Player struct {
	Name  string
	Score int
}

func NewPlayer() *Player {
	p := Player{
		Name: fmt.Sprintf("%s Player", randomEmoji()),
	}
	return &p
}

func (p *Player) AddScore() {
	fmt.Printf("%s got a score!\n", p.Name)
	p.Score++
}

func (p *Player) Play(ask chan *quiz.Question, answer chan *quiz.Answer) {
	question := <-ask

	//think
	time.Sleep(time.Second * time.Duration(rand.Int31n(10)))

	selectedOption := rand.Int31n(int32(len(question.Options)))
	fmt.Printf("%s selected option %d\n", p.Name, selectedOption)
	selected := selectedOption
	answer <- &quiz.Answer{
		Caller:         p.Name,
		SelectedOption: int(selected),
	}
}

func randomEmoji() string {
	// http://apps.timwhitlock.info/emoji/tables/unicode
	emoji := [][]int{
		// Emoticons icons
		{128513, 128591},
		// Transport and map symbols
		{128640, 128704},
	}
	r := emoji[rand.Int()%len(emoji)]
	min := r[0]
	max := r[1]
	n := rand.Intn(max-min+1) + min
	return html.UnescapeString("&#" + strconv.Itoa(n) + ";")
}
