package main

import (
	"fmt"

	"github.com/josh-dev-07/physics-lib/pkg/physics"
)

func main() {
	mass1 := 1000.0  // Mass of object 1 in kilograms
	mass2 := 500.0   // Mass of object 2 in kilograms
	distance := 10.0 // Distance in meters

	force := physics.GravitationalForce(mass1, mass2, distance)
	fmt.Printf("Gravitational Force: %f Newtons\n", force)
}
