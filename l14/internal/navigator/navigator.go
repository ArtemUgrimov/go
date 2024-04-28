package navigator

import (
	"hw15/internal/primitives"
	"math/rand"
)

// Business layer
func CalculateRoute(start primitives.RoutePoint, finish primitives.RoutePoint) []primitives.RoutePoint {
	result := make([]primitives.RoutePoint, 0)
	for range rand.Intn(5) + 2 { //minimum 2 points
		result = append(result, primitives.RoutePoint{
			Latitude:   rand.Float64(),
			Longtitude: rand.Float64(),
		})
	}
	return result
}
