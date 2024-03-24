package routs

import "main/l6/transports"

type RouteStage struct {
	Transport transports.PublicTransport
	Distance  int32
	Direction string
}
