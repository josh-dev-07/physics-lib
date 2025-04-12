# 🔭 Physics Library

**physics-lib** is a Go library that provides a collection of common physics formulas—ideal for A-level topics including mechanics and gravitational theory. Use this library to calculate values for kinetic energy, potential energy, escape velocity, gravitational force, momentum, pressure, acceleration, and more.

## Table of Contents

- [🔭 Physics Library](#-physics-library)
  - [Table of Contents](#table-of-contents)
  - [Installation](#installation)
  - [Usage](#usage)
  - [Examples](#examples)
    - [Calculate Potential Energy](#calculate-potential-energy)
    - [Calculate Escape Velocity](#calculate-escape-velocity)
    - [Calculate Gravitational Force](#calculate-gravitational-force)
  - [Available Formulas](#available-formulas)
  - [Constants](#constants)
  - [Contributions](#contributions)

## Installation

To add **physics-lib** to your Go project, use `go get`:

```bash
go get github.com/josh-dev-07/physics-lib
```

## Usage

Import the library in your Go file:

```go
import (
	"fmt"
	"github.com/josh-dev-07/physics-lib/pkg/physics"
)
```

## Examples

### Calculate Potential Energy

```go
mass := 10.0   // in kilograms
height := 5.0  // in meters

pe := physics.PotentialEnergy(mass, height)
fmt.Printf("Potential Energy: %f Joules\n", pe)
```

### Calculate Escape Velocity

```go
planetMass := 5.972e24     // Earth's mass in kilograms
planetRadius := 6.371e6    // Earth's radius in meters

ev := physics.EscapeVelocity(planetMass, planetRadius)
fmt.Printf("Escape Velocity: %f m/s\n", ev)
```

### Calculate Gravitational Force

```go
mass1 := 1000.0      // Mass of object 1 in kilograms
mass2 := 500.0       // Mass of object 2 in kilograms
distance := 10.0     // Distance in meters

force := physics.GravitationalForce(mass1, mass2, distance)
fmt.Printf("Gravitational Force: %f Newtons\n", force)
```

## Available Formulas

- Kinetic Energy: ½ * m * v²
- Potential Energy: m * g * h
- Escape Velocity: v = √((2 * G * M) / R)
- Escape Energy: ½ * m * (Escape Velocity)²
- Gravitational Force: F = G * (m₁ * m₂) / r²
- Momentum: p = m * v
- Pressure: P = F / A
- Acceleration: a = (v_final - v_initial) / t

## Constants

Available physical constants in `consts` package:

- `GravitationalConstant = 6.67430e-11` (m³·kg⁻¹·s⁻²)
- `StandardGravity = 9.80665` (m/s²)

## Contributions

Contributions are welcome! Please submit pull requests or issues if you have improvements or find issues.
