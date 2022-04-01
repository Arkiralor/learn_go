//Program to play "BingBong"

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var i int
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the upper limit: ")
	upper_limit_str, _ := reader.ReadString('\n')
	upper_limit_str = strings.Replace(upper_limit_str, "\n", "", -1)
	upper_limit, err := strconv.Atoi(upper_limit_str)
	fmt.Print("Error Code: ", err)
	num1 := 3
	num2 := 5
	for i = 1; i <= upper_limit; i += 1 {
		if i%num1 == 0 && i%num2 != 0 {
			fmt.Print("\nBing: ", i)
		} else if i%num1 != 0 && i%num2 == 0 {
			fmt.Print("\nBong: ", i)
		} else if i%num1 == 0 && i%num2 == 0 {
			fmt.Print("\nBingBong: ", i)
		} else {
			fmt.Print("\nInvalid value: ", i)
		}
	}
	fmt.Println("\n")

}
