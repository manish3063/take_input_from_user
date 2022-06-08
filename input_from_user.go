// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	fmt.Println("Enter 1st number")
	var num1 int
	fmt.Scanln(&num1)
	fmt.Println("Enter 2nd number")
	var num2 int
	fmt.Scanln(&num2)

	sum := num1 + num2
	fmt.Println("sum is :", sum)
}
