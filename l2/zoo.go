package main

import "fmt"

type Zookeeper struct {
	Name string
}

type Cage struct {
	Animal Cagable
}

func (c *Cage) Occupied() bool {
	return c.Animal != nil
}

func (c *Cage) Occupy(animal Cagable) bool {
	if animal.isFree() && !c.Occupied() {
		c.Animal = animal
		animal.imprison()
		return true
	}
	return false
}

func (c *Cage) freeAnimal() bool {
	if c.Occupied() {
		c.Animal.free()
		c.Animal = nil
		return true
	}
	return false
}

func catchTheAnimal(zk *Zookeeper, a Cagable, c *Cage) {
	if a.isFree() {
		fmt.Printf("%s caught the %s\n", zk.Name, a.whoAreYou())
		c.Occupy(a)
	} else {
		fmt.Printf("%s cannot put %s into the cage\n", zk.Name, a.whoAreYou())
	}
}
