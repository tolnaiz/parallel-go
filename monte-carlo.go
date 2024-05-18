package main

import "fmt"
import "math/rand"

func main() {
	var npoints = 1000000
	var circle_count = 0
	var pi = 0.0
	for i := 0; i < npoints; i++ {
		var x = rand.Float64()
		var y = rand.Float64()
		if (x*x + y*y) <= 1 {
			circle_count++
		}
		pi = 4.0 * float64(circle_count) / float64(npoints)
		fmt.Printf("%f at iteration %d \r", pi, i)
	}
	fmt.Print("\n")
}