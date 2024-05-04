package routes

import (
	"errors"
	"math/rand"

	"hw15/internal/navigator"
	"hw15/internal/primitives"
)

// Application layer
type RouteStatus int8

const (
	Planned RouteStatus = iota
	InProgress
	Finished
)

type Route struct {
	Id              int64
	Status          RouteStatus
	Points          []primitives.RoutePoint
	NextDestination int
}

func NewRoute(start primitives.RoutePoint, end primitives.RoutePoint) Route {
	route := Route{
		Id:     rand.Int63(),
		Status: Planned,
	}
	route.Points = navigator.CalculateRoute(start, end)
	route.NextDestination = 0
	return route
}

func (r *Route) Start() error {
	if r.Status != Planned {
		return errors.New("cannot start route")
	}
	r.Status = InProgress
	return r.Next()
}

func (r *Route) Next() error {
	if r.NextDestination < len(r.Points) || r.Status == InProgress {
		r.NextDestination++
		if r.NextDestination == len(r.Points) {
			r.Status = Finished
		}
		return nil
	} else {
		return errors.New("cannot proceed to the next point")
	}
}

func (r *Route) IsFinished() bool {
	return r.Status == Finished
}
