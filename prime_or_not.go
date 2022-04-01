//Program to find whether a given number is prime or not.package learn_go

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check_prime(num int) bool {
	i := 2
	var flag_var bool = false
	var remainder int
	var upper_limit int = num / 2

	for i <= upper_limit {
		remainder = num % i
		if remainder == 0 {
			flag_var = true
		}
		i += 1
	}

	return flag_var

}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the highest value till which we have to check: ")
	upper_limit_str, _ := reader.ReadString('\n')
	upper_limit_str = strings.Replace(upper_limit_str, "\n", "", -1)
	upper_limit, _ := strconv.Atoi(upper_limit_str)

	for i := 3; i <= upper_limit; i += 1 {
		is_prime := check_prime(i)
		if is_prime == false {
			fmt.Println(i, " is a Prime Number.")
		}
	}

	fmt.Print("\n")

}
