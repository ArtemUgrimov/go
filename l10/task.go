package main

type Task struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	IsCompleted bool   `json:"is_completed"`
}
