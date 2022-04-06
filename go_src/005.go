//Program using package-level (global) variable

package main

import (
	"fmt"
)

var a int = 9

func demo() {
	a := 8
	fmt.Println("In Demo(): a =", a)
}

func main() {
	demo()
	fmt.Println("In Main(): a =", a)
}
