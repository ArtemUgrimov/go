package transports

import (
	"main/l6/passengers"
)

type Bus struct {
	passengers []*passengers.Passenger
}

func (b *Bus) PickUpPassenger(pass *passengers.Passenger) {
	b.passengers = append(b.passengers, pass)
}

func (b *Bus) DropPassenger() {
	if b.passengers == nil {
		return
	}
	ln := len(b.passengers)
	if ln == 0 {
		return
	}
	b.passengers = b.passengers[:ln-1]
}

func (b *Bus) ToString() string {
	return "bus"
}

type Train struct {
	passengers []*passengers.Passenger
}

func (b *Train) PickUpPassenger(pass *passengers.Passenger) {
	b.passengers = append(b.passengers, pass)
}

func (b *Train) DropPassenger() {
	if b.passengers == nil {
		return
	}
	ln := len(b.passengers)
	if ln == 0 {
		return
	}
	b.passengers = b.passengers[:ln-1]
}

func (t *Train) ToString() string {
	return "train"
}

type Airplane struct {
	passengers []*passengers.Passenger
}

func (b *Airplane) PickUpPassenger(pass *passengers.Passenger) {
	b.passengers = append(b.passengers, pass)
}

func (b *Airplane) DropPassenger() {
	if b.passengers == nil {
		return
	}
	ln := len(b.passengers)
	if ln == 0 {
		return
	}
	b.passengers = b.passengers[:ln-1]
}

func (a *Airplane) ToString() string {
	return "airplane"
}
