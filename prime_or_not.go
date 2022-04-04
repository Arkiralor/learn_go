//Program to find whether a given number is prime or not.package learn_go

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func check_prime(num int) bool {
	i := 2
	flag_var := true
	var remainder int
	var upper_limit int = num / 2

	for i <= upper_limit {
		remainder = num % i
		if remainder == 0 {
			flag_var = false
		}
		i += 1
	}

	return flag_var

}

func read_number_from_stdin() int {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the highest value till which we have to check: ")
	upper_limit_str, _ := reader.ReadString('\n')
	upper_limit_str = strings.Replace(upper_limit_str, "\r", "", -1)
	upper_limit_str = strings.Replace(upper_limit_str, "\n", "", -1)
	upper_limit, err := strconv.Atoi(upper_limit_str)

	if err != nil {
		fmt.Println("Error Code: ", err)
		os.Exit(1)
	}

	return upper_limit
}

func main() {
	var i int
	var prime_number []int = []int{2}

	upper_limit := read_number_from_stdin()

	t1 := time.Now()
	for i = 3; i <= upper_limit; i += 1 {
		fmt.Println("Checking: ", i)
		is_prime := check_prime(i)
		if is_prime == true {
			prime_number = append(prime_number, i)
		}
	}

	t2 := time.Now()
	fmt.Printf("\nThe list of prime numbers until %d is: %v\n", upper_limit, prime_number)

	fmt.Printf("This operation took %v.\n", t2.Sub(t1))

}
