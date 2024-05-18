package main

import "fmt"
import "math"
import "math/rand"
import "runtime"

type point struct {
    x float64
    y float64
}


func main() {
	var npoints = 10000
	var all_circle_count = 0
	var pi = 0.0
	var cpus = runtime.NumCPU()
	var iterations = npoints / cpus
	count_channel := make(chan int)

	fmt.Printf("Using %d CPUs\n", cpus)

	for j := 0; j < cpus; j++ {
		go func() {
			var circle_count = 0
			for i := 0; i < iterations; i++ {
				var point = point{rand.Float64(), rand.Float64()}
				if (point.x * point.x + point.y * point.y) <= 1 {
					circle_count++
				}
			}
			count_channel <- circle_count
		}()
	}
	for j := 0; j < cpus; j++ {
		all_circle_count += <-count_channel
		pi = 4.0 * float64(all_circle_count) / float64(iterations * (j + 1))
		fmt.Printf("%.10f (error %.10f) at iteration %d \n", pi, math.Pi - pi, iterations * (j + 1))
	}
	fmt.Print("\n")
}