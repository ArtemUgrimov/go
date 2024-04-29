package primitives

import "fmt"

type RoutePoint struct {
	Latitude   float64
	Longtitude float64
}

func (r *RoutePoint) String() string {
	return fmt.Sprintf("%f, %f", r.Latitude, r.Longtitude)
}
