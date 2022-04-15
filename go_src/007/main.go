//Program to write to a file

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func read_number_from_stdin() float64 {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the highest value till which we have to check: ")
	num_str, _ := reader.ReadString('\n')
	num_str = strings.Replace(num_str, "\r", "", -1)
	num_str = strings.Replace(num_str, "\n", "", -1)
	num, err := strconv.Atoi(num_str)

	if err != nil {
		fmt.Println("Error Code: ", err)
		os.Exit(1)
	}

	return float64(num)
}

func main() {
	fmt.Println("Writing to file.")

	str1 := "This is only a test."
	data := []byte(str1)

	file, err := os.Create("data_files/001.txt")

	if err != nil {
		fmt.Println("Error Code: ", err)
		os.Exit(1)
	}
	file.Close()

	write_err := os.WriteFile("data_files/001.txt", data, 0644)
	if write_err != nil {
		fmt.Println("Error Code: ", write_err)
		os.Exit(1)
	}

}
