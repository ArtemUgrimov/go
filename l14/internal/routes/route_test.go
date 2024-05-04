package routes

import (
	"hw15/internal/primitives"
	"testing"

	"github.com/stretchr/testify/require"
)

func NewTestRoute() Route {
	route := NewRoute(primitives.RoutePoint{
		Latitude:   123.0,
		Longtitude: 456.0,
	}, primitives.RoutePoint{
		Latitude:   123.0,
		Longtitude: 456.0,
	})
	return route
}

func TestRouteCreation(t *testing.T) {
	route := NewTestRoute()
	// check that route has status "Planned"
	require.Equal(t, Planned, route.Status)
	// check if the route has more than two points after creation
	require.Condition(t, func() (success bool) {
		success = len(route.Points) >= 2
		return
	}, "must be two or more points")
}

func TestRouteProcessingPositive(t *testing.T) {
	// route becomes finished after len(points) steps
	route := NewTestRoute()
	err := route.Start()
	require.NoError(t, err)

	for range len(route.Points) - 2 {
		err = route.Next()
		require.NoError(t, err)
	}
	err = route.Next()
	require.NoError(t, err)
	require.Equal(t, Finished, route.Status)
}

func TestRouteProcessingNegative(t *testing.T) {
	// route gives an error after > len(points) steps
	route := NewTestRoute()
	err := route.Start()
	require.NoError(t, err)

	for range len(route.Points) - 1 {
		err = route.Next()
		require.NoError(t, err)
	}
	err = route.Next()
	require.Error(t, err)
}
