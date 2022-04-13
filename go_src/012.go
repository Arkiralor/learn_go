//Program to find factors of a given number.

package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"math/rand"
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

func find_factors(num int) []int {
	var factors []int = []int{1}
	var upper_limit int = num / 2

	for i := 2; i <= upper_limit; i += 1 {
		if num%i == 0 {
			to_append := convert_to_binary(int64(i))
			fmt.Println("Factor found: ", to_append)
			factors = append(factors, to_append)
		}
	}

	return factors
}

func convert_to_binary(num int64) int {
	num_str := strconv.FormatInt(num, 2)
	num_int, err := strconv.Atoi(num_str)
	if err != nil {
		log.Fatalf("\nError Code: %v\n", err.Error())
	}
	return num_int
}

func random_binary_number(size int) int {
	rand_bin, current_size := 0, 0
	var rand_bit int
	for current_size < size {
		rand_bit = rand.Intn(2)
		rand_bin = rand_bin*10 + rand_bit
		current_size += 1
	}
	log.Printf("\nRandom binary number: %d\n", rand_bin)
	return rand_bin
}

func binary_to_int(binary_num int) int {
	var decimal_num int = 0
	var power int = 0
	for binary_num > 0 {
		remainder := binary_num % 10
		decimal_num += remainder * int(math.Pow(2, float64(power)))
		binary_num /= 10
		power += 1
	}
	log.Printf("\nConverted %d to %d.\n", binary_num, decimal_num)
	return decimal_num
}

func main() {
	fmt.Print("Enter the number of bits to generate: ")
	var bits int = read_number_from_stdin()
	var num int = binary_to_int(random_binary_number(bits))

	t1 := time.Now()
	factors := find_factors(num)

	t2 := time.Now()
	fmt.Printf("\nThe list of factors of %d (%d) is: %v\n", convert_to_binary(int64(num)), num, factors)

	fmt.Printf("\nThis operation took %v.\n", t2.Sub(t1))
}
