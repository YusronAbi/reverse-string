package main

import "fmt"

func reverseString(input string) string {
	reversed := ""
	for _, char := range input {
		reversed = string(char) + reversed
	}
	return reversed
}

func main() {
	fmt.Println(reverseString("hello"))
}
