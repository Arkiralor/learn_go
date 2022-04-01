// Go Program to read a user's name and print a "Hello world" line using the name.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name: ")
	text, _ := reader.ReadString('\n')
	fmt.Println("Hello, ",text)
}
