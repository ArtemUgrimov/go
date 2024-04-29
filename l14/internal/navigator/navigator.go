package navigator

import (
	"hw15/internal/primitives"
	"math/rand"
)

// Business layer
func CalculateRoute(start primitives.RoutePoint, finish primitives.RoutePoint) []primitives.RoutePoint {
	result := make([]primitives.RoutePoint, 0)
	result = append(result, start)
	for range rand.Intn(5) {
		result = append(result, primitives.RoutePoint{
			Latitude:   rand.Float64(),
			Longtitude: rand.Float64(),
		})
	}
	result = append(result, finish)
	return result
}
