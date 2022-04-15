package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Function to take a string input from the console and convert it into a number (int).
func NumberFromCon() int {
	reader := bufio.NewReader(os.Stdin)
	input_str, inp_err := reader.ReadString('\n')
	if inp_err != nil {
		log.Fatalf("Error Code: %v", inp_err)
	}
	input_str = RemoveEOLMarkers(input_str)
	input_int, err := strconv.Atoi(input_str)

	if err != nil {
		log.Fatalf("Error Code: %v", err)
	}

	return input_int
}

// Function to create a list from a given string, where the separation occurs at the whitespace.
func ListFromString(s string) []string {
	s = RemoveEOLMarkers(s)
	s_2 := strings.Split(s, " ")
	return s_2
}

// Function to remove EOL markers (\n in Unix, \n\r in Windows) from a given string.
func RemoveEOLMarkers(s string) string {
	s = strings.Replace(s, "\r", "", -1)
	s = strings.Replace(s, "\n", "", -1)

	return s
}

// Function to remove white spaces from a given string.
func RemoveWhiteSpaces(s string) string {
	s = strings.Replace(s, " ", "", -1)

	return s
}

func main() {
	arr := "Oh what a travesty it is to be in love!"
	fmt.Println(arr)
	new_var := ListFromString(arr)
	fmt.Println(new_var)

}
