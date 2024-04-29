package navigator

import (
	"hw15/internal/primitives"
	"testing"
)

func TestNavigatorPositiveLen(t *testing.T) {
	start := primitives.RoutePoint{
		Latitude:   123.0,
		Longtitude: 456.0,
	}
	finish := primitives.RoutePoint{
		Latitude:   0.0,
		Longtitude: 9000.0,
	}
	route := CalculateRoute(start, finish)
	if len(route) < 2 {
		t.Fatalf("Route cannot be with less than 2 points")
	}
}

func TestNavigatorStartFinish(t *testing.T) {
	start := primitives.RoutePoint{
		Latitude:   123.0,
		Longtitude: 456.0,
	}
	finish := primitives.RoutePoint{
		Latitude:   0.0,
		Longtitude: 9000.0,
	}
	route := CalculateRoute(start, finish)
	if route[0].Latitude != start.Latitude || route[0].Longtitude != start.Longtitude {
		t.Fatalf("Start is different than expected %s vs %s", start.String(), route[0].String())
	}
	if route[len(route)-1].Latitude != finish.Latitude || route[len(route)-1].Longtitude != finish.Longtitude {
		t.Fatalf("Start is different than expected %s vs %s", start.String(), route[0].String())
	}
}
