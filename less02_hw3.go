package main

import "fmt"

func main() {

	var (
		x float64 = 5.3
		y float64 = 2.696
	)

	fmt.Printf("Initial state\nx: %.3f, y: %.3f\n\n", x, y)
	// Initial state
	// x: 5.300, y: 2.696

	// pass the correct arguments to setCenter func.
	setCenter(nil, nil)

	fmt.Printf("Centered!\nx: %.1f, y: %.1f\n\n", x, y)
	// Centered!
	//x: 0.0, y: 0.0

}

func setCenter(a, b **float64) {
	// change the function body in a such way
	// that it set original coordinates x and y to 0.
	panic("implement me!")
}
