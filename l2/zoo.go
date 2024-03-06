package main

import "fmt"

type Zookeeper struct {
	Name string
}

type Cage struct {
	Animal   *Animal
	Occupied bool
}

func (c *Cage) putAnimal(a *Animal) bool {
	if !c.Occupied && a.IsFree {
		c.Animal = a
		c.Occupied = true
		a.IsFree = false
		return true
	}
	return false
}

func (c *Cage) freeAnimal() bool {
	if c.Occupied {
		c.Animal.IsFree = true
		c.Animal = nil
		c.Occupied = false
		return true
	}
	return false
}

func catchTheAnimal(zk *Zookeeper, a *Animal, c *Cage) {
	if c.putAnimal(a) {
		fmt.Printf("%s caught the %s\n", zk.Name, a.whoAreYou())
	} else {
		fmt.Printf("%s cannot put %s into the cage\n", zk.Name, a.whoAreYou())
	}
}
