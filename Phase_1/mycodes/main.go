package main  
import "fmt"

/*example of basic function

func add(num1 int, num2 int) {    //same logic everywhere, type after var name

	sum := num1 + num2

	fmt.Println("Sum is:", sum)

} */

func main() {
	fmt.Println("hello, world") // println will print in a line and go to the next line

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
    var x, y, z = 10, "This is string", 30 // here values are assigned serially
	fmt.Println(x, y,z) // prints everything in the same line and goes to next, auto type inference

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
