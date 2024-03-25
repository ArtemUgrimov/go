package routs

import (
	"fmt"
	"main/l6/passengers"
)

type Route struct {
	stages []RouteStage
}

func (r *Route) AddTransport(description RouteStage) {
	if r.stages == nil {
		r.stages = make([]RouteStage, 0)
	}
	r.stages = append(r.stages, description)
}

func (r *Route) Navigate(pass *passengers.Passenger) {
	if len(r.stages) == 0 {
		fmt.Println("You are at the destination")
		return
	}
	for index, stage := range r.stages {
		fmt.Printf("Go to the %s by a the %s", stage.Direction, stage.Transport.ToString())
		stage.Transport.PickUpPassenger(pass)
		if index < len(r.stages)-1 {
			fmt.Println(" and then...")
			stage.Transport.DropPassenger()
		} else {
			fmt.Println(" and you are at the destination")
		}
	}
}
