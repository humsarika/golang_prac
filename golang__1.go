//Variables in golang
package main

import "fmt"

func main() {
	//there are two ways to declare a variable in golang
	//1.using var keyword
	var a = 10
	var b int
	//1.using short variable declaration
	c := "sarika kushwaha"

	//taking input from user
	fmt.Scan(&b)

	//displaying value of variables in golang
	fmt.Printf("Value of a = %v \nb =%v \n c =%v\n", a, b, c)
	//displaying type of variables in golang
	fmt.Printf("Type of a = %T and c =%T\n", a, c)
	//displaying type of variables in golang
	fmt.Printf("Address of a = %v\nAddress of b = %v\nAddress of c =%v ", &a, &b, &c)

}
