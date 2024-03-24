package main

import (
	"main/l6/passengers"
	"main/l6/routs"
	"main/l6/transports"
)

func main() {
	route := routs.Route{}
	route.AddTransport(routs.RouteStage{
		Transport: &transports.Bus{},
		Distance:  100,
		Direction: "West",
	})
	route.AddTransport(routs.RouteStage{
		Transport: &transports.Airplane{},
		Distance:  2300,
		Direction: "East-North",
	})
	route.AddTransport(routs.RouteStage{
		Transport: &transports.Train{},
		Distance:  200,
		Direction: "South",
	})

	artem := passengers.Passenger{
		Weight: 78,
	}
	route.Navigate(&artem)
}
