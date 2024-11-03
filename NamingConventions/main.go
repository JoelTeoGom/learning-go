package main

import (
	"errors"
	"fmt"
	"math"
)

// Constants
const PI = 3.14159
const maxUsers = 100

// Error constants
var ErrNotFound = errors.New("item not found")

// Shape interface following the -er naming convention
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Circle struct as an example of a concrete type implementing Shape
type Circle struct {
	Radius float64
}

// Implementing the Shape interface for Circle
func (c Circle) Area() float64 {
	return PI * math.Pow(c.Radius, 2)
}

func (c Circle) Perimeter() float64 {
	return 2 * PI * c.Radius
}

// Function following camelCase naming
func calculateTotal(radius float64) float64 {
	circle := Circle{Radius: radius}
	return circle.Area() + circle.Perimeter()
}

func main() {
	fmt.Println("Total calculation for Circle:", calculateTotal(5))

	// Error handling example
	if err := findItem("nonexistent"); err != nil {
		fmt.Println("Error:", err)
	}
}

// Simulating an error return for demonstration
func findItem(item string) error {
	return ErrNotFound
}
