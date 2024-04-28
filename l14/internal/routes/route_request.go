package routes

import "hw15/internal/primitives"

type RouteRequest struct {
	Start  primitives.RoutePoint `json:"start"`
	Finish primitives.RoutePoint `json:"finish"`
}
