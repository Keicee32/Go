package main

import "fmt"

func main() {
	fmt.Print("Enter a number: ")
	firstNumber := 0
	fmt.Scanln(&firstNumber)

	bin := fmt.Sprintf("The value in binary is " + "%b" , firstNumber)
	fmt.Println(bin)
}