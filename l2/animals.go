package main

import "fmt"

type Cagable interface {
	isFree() bool
	free()
	imprison()
	whoAreYou() string
}

type Animal struct {
	Name   string
	IsFree bool
}

func (a *Animal) isFree() bool {
	return a.IsFree
}

func (a *Animal) free() {
	a.IsFree = true
}

func (a *Animal) imprison() {
	a.IsFree = false
}

func (a *Animal) whoAreYou() string {
	if a.IsFree {
		return fmt.Sprintf("Free %s", a.Name)
	} else {
		return fmt.Sprintf("Imprisoned %s", a.Name)
	}
}

type Monkey struct {
	Animal
	JumpHeight uint32
}

type Giraffe struct {
	Animal
	NeckLength uint32
}

type Lion struct {
	Animal
	Roars bool
}
