//Loops in Golang

package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 5; i += 1 {
		fmt.Println("This is a loop; iteration:", i)
	}

	for j := 1; j <= 100; j += 1 {
		fmt.Print(j)
		if j != 100 {
			fmt.Print(", ")
		}
		if j%10 == 0 && j != 100 {
			fmt.Print("\n")
		} else if j == 100 {
			fmt.Print(". ")
		}
	}
	fmt.Print("\n")
}
