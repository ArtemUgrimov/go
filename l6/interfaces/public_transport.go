package interfaces

import (
	"main/l6/passengers"
)

type PublicTransport interface {
	PickUpPassenger(pass *passengers.Passenger)
	DropPassenger()
	ToString() string
}
