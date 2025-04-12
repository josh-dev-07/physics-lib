package main

import (
	"fmt"

	"github.com/josh-dev-07/physics-lib/pkg/physics"
)

func main() {
	planetMass := 5.972e24  // Earth's mass in kilograms
	planetRadius := 6.371e6 // Earth's radius in meters

	ev := physics.EscapeVelocity(planetMass, planetRadius)
	fmt.Printf("Escape Velocity: %f m/s\n", ev)
}
