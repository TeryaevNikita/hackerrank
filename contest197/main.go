package main

import (
	"fmt"
	"math"
)

func main() {
	var un, b float64 = 0, 30
	var un1 float64

	ff := func(un, b float64) float64 {
		power := b - un*un
		return math.Floor(math.Pow(2, power)) * 1E-9
	}

	for i := float64(0); i <= 20E1; i++ {
		un1 = un
		un = ff(un1, b)
		fmt.Println(un)
		if math.Round(un1*1E12) == math.Round(un*1E12) {
			break
		}
	}

	fmt.Printf("%.9f\n", un1+un)
}
