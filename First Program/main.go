package main

import "fmt"

/*example of basic function

func add(num1 int, num2 int) {    //same logic everywhere, type after var name

	sum := num1 + num2

	fmt.Println("Sum is:", sum)

} */

func main() {
	fmt.Println("hello, world")

	/*

		all go files will start with a package declaration

		package main means executable program, other packages are library

		main package will contain func main

	*/

	//x := 20
	//y := 30

	//add(x, y)

	//-----------------------------------------------------------------------------------------------------------------------------------------------------
	/* conditionals (if else, switch)  && || ! they are same in go

	a := 18

	if a > 18 {
		fmt.Println("Eligible")
	} else if a < 18 {                          //next block same line after }
		fmt.Println("Ineligible")
	} else {
		fmt.Println("Just wait a little bit")
	}

	isPretty := false

	if !isPretty {
		fmt.Println("conditional block executes when true")
	} */

	//-----------------------------------------------------------------------------------------------------------------------------------------------
	/*conditionals switch
	    t := 12
		switch t {
		case 1:
			fmt.Println("Not found")

		case 10, 12:
			fmt.Println("Found")

		default:
			fmt.Println("Bye Bye")
		}
	*/

	//---------------------------------------------------------------------------------------------------------------------------------------------------

}
