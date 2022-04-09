package main

import (
	"fmt"
)

type Dict struct {
	key   string
	value string
}

func main() {

	var a Dict
	a = Dict{
		key:   "key1",
		value: "value1",
	}
	a = Dict{
		key:   "key2",
		value: "value2",
	}

	fmt.Println(a)
}
