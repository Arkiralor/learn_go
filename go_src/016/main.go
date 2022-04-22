package main

import "fmt"

func pow(base float64, exp int64) float64 {
	var ret float64
	if exp > 0 {
		if exp == 1 {
			ret = base
		} else {
			ret = base * pow(base, exp-1)
		}
	} else if exp < 0 {
		if exp == -1 {
			ret = 1 / base
		} else {
			ret = 1 / (base / pow(base, exp+1))
		}
	} else {
		ret = 1
	}
	return ret
}

func main() {
	res := pow(2, -5)
	fmt.Println(res)
}
