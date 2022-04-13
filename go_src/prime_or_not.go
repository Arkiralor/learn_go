//Program to find whether a given number is prime or not.package learn_go

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
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
	input_str, _ := reader.ReadString('\n')
	input_str = clean_up_stdin(input_str)

	input_int, err := strconv.Atoi(input_str)

	if err != nil {
		fmt.Println("Error Code: ", err)
		os.Exit(1)
	}

	return input_int
}

func clean_up_stdin(s string) string {
	s = strings.Replace(s, "\r", "", -1)
	s = strings.Replace(s, "\n", "", -1)
	s = strings.Replace(s, " ", "", -1)
	return s
}

func main() {
	var i int
	var prime_number []int = []int{2}
	fmt.Print("Enter the highest value till which we have to check: ")
	var upper_limit int = read_number_from_stdin()

	t1 := time.Now()
	for i = 3; i <= upper_limit; i += 1 {
		is_prime := check_prime(i)
		if is_prime {
			log.Printf("Prime number found: %d", i)
			prime_number = append(prime_number, i)
		}
	}

	t2 := time.Now()
	fmt.Printf("\nThe list of prime numbers until %d is: %v\n", upper_limit, prime_number)

	fmt.Println("The type of the time difference is: ", reflect.TypeOf(t2.Sub(t1)))

	fmt.Printf("\nThis operation took %v.\n", t2.Sub(t1))
}
