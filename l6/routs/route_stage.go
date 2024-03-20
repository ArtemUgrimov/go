package routs

import "main/l6/interfaces"

type RouteStage struct {
	TransportType interfaces.PublicTransport
	Distance      int32
	Direction     string
}
