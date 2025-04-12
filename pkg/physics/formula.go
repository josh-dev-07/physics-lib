package physics

import (
	"math"

	"github.com/josh-dev-07/physics-lib/pkg/consts"
)

// PotentialEnergy returns m * g * h.
func PotentialEnergy(mass, height float64) float64 {
	return mass * consts.StandardGravity * height
}

// EscapeVelocity returns the speed needed to leave a planet's gravitational field. v = √((2 * G * M) / R)
func EscapeVelocity(planetMass, planetRadius float64) float64 {
	return math.Sqrt((2 * consts.GravitationalConstant * planetMass) / planetRadius)
}

// KineticEnergy returns ½ * m * v².
func KineticEnergy(mass, velocity float64) float64 {
	return 0.5 * mass * math.Pow(velocity, 2)
}

// EscapeEnergy returns the energy required for an object to escape the gravitational field.
func EscapeEnergy(mass, planetMass, planetRadius float64) float64 {
	ev := EscapeVelocity(planetMass, planetRadius)
	return 0.5 * mass * math.Pow(ev, 2)
}

// GravitationalForce computes the gravitational force between two masses.
// Formula: F = G * (m1*m2) / r².
func GravitationalForce(m1, m2, distance float64) float64 {
	return consts.GravitationalConstant * (m1 * m2) / math.Pow(distance, 2)
}

// Momentum calculates the momentum of an object.
// Formula: p = m * v.
func Momentum(mass, velocity float64) float64 {
	return mass * velocity
}

// Pressure calculates the pressure from a force distributed over an area.
// Formula: P = F / A.
func Pressure(force, area float64) float64 {
	return force / area
}

// Acceleration calculates the average acceleration given a change in velocity over time.
// Formula: a = (v_final - v_initial) / t.
func Acceleration(initialVelocity, finalVelocity, time float64) float64 {
	return (finalVelocity - initialVelocity) / time
}
