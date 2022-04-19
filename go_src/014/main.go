package main

import (
	"fmt"
	"math"
)

func ConvertDegToRadian(angle float64) float64 {
	return angle * (math.Pi / 180)
}

func main() {
	var angle float64 = 90.00

	angle = ConvertDegToRadian(angle)

	fmt.Println(math.Sin(angle))
}
