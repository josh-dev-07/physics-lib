package main

import (
	"fmt"

	"github.com/josh-dev-07/physics-lib/pkg/physics"
)

func main() {
	mass := 10.0  // in kilograms
	height := 5.0 // in meters

	pe := physics.PotentialEnergy(mass, height)
	fmt.Printf("Potential Energy: %f Joules\n", pe)
}
