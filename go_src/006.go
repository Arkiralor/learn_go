//Arrays and Slices

package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Arrays
	var ages [3]int = [3]int{23, 34, 45}
	names := [3]string{"John", "Paul", "George"}

	fmt.Println(ages, reflect.TypeOf(ages))
	fmt.Println(names, len(names), reflect.TypeOf(names))

	// Slices - dynamic array (use array as a base)
	var names2 []string = []string{"John", "Paul", "George", "Ringo", "Pete", "Vasquez"}
	fmt.Println(names2, len(names2), reflect.TypeOf(names2))

	names2 = append(names2, "Santana", "Jules")
	fmt.Println(names2, len(names2), reflect.TypeOf(names2))
	fmt.Println(names2[2-1:5+2], reflect.TypeOf(names2[2-1:5+2]))

}
