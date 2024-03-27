package game

import (
	"fmt"
	"main/l8/quiz"
)

type Round struct {
	Question    quiz.Question `json:"Question"`
	RightAnswer int           `json:"RightAnswer"`
}

type IPlayer interface {
	Play(question chan *quiz.Question, answer chan *quiz.Answer)
	AddScore()
}

func NewRound(question string, answers []string, rightAnswer int) *Round {
	return &Round{
		Question: quiz.Question{
			Text:    question,
			Options: answers,
		},
		RightAnswer: rightAnswer,
	}
}

func (r *Round) Run(questions chan *quiz.Question, amount int) {
	fmt.Println("===> Running a round : <===")
	fmt.Println("  ", r.Question.Text)
	for idx, opt := range r.Question.Options {
		fmt.Println("     ", idx, opt)
	}

	for i := 0; i < amount; i++ {
		questions <- &r.Question
	}
}

func (r *Round) Check(answers chan *quiz.Answer, players map[string]IPlayer) {
	for i := 0; i < len(players); i++ {
		answer := <-answers
		if r.RightAnswer == answer.SelectedOption {
			players[answer.Caller].AddScore()
		}
	}
}
