package main

import "fmt"
import "math"
import "math/rand"
import "os"
import "strconv"

func main() {
	var npoints = 1000000
	var err error
	if len(os.Args) > 1 {
		npoints, err = strconv.Atoi(os.Args[1])
		if err != nil {
			panic(err)
		}
	}

	var circle_count = 0
	var pi = 0.0
	for i := 0; i < npoints; i++ {
		var x = rand.Float64()
		var y = rand.Float64()
		if (x*x + y*y) <= 1 {
			circle_count++
		}
		pi = 4.0 * float64(circle_count) / float64(i)
		fmt.Printf("%.10f (error %.10f) at iteration %d \r", pi, math.Pi - pi, i)
	}
	fmt.Print("\n")
}