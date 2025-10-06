package main

import "fmt"

func main() {
	var x, y, z = 10, "This is string", 30
	fmt.Println(x, y, z) // prints everything in the same line and goes to next, auto type inference

	var a string = "type after"
	fmt.Println(a) // if declared then afer the variable name

	// short form , can use only inside function
	age := 21
	fmt.Println(age)

	//no assignment must use var keyword , also must declare type for go to infer it

	var time int
	fmt.Println(time) // no value then prints 0 for float

	var khaisi bool
	fmt.Println(khaisi) // prints flase for no assignment

	// re-assign e colon deya jabena

	pp := 12
	fmt.Println(pp)
	pp = 34 // re-assigned colon hobena, type same hote hobe

	fmt.Println(pp)

}
