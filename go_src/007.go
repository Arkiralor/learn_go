//Program to write to a file

package main

import (
	"fmt"
	"os"
)

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
