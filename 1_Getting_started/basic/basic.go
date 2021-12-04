/*
	GO IS a programming language
	to improve programming productivity in an
	era of multicore, networked machines and large codebases

*/

package main

/*
	every go program is made up of packages
	go program start running in --main-- package,
	this is why we need to declare our code as the main package
	to make it run and create the output

	PACKAGES: many package that can be imported and used in
	the code to accomplish different task
	fmt stands for format, I/O functionality

	import math
*/

import (
	"fmt"
)

/*
	in Go, a name is exported if it begins with a capital letter
can access the exported names using the package name, a dot,

*/

func main() {
	/*
		prior to other programming language
		is the entry point of our program,
		it's the function taht gets execute when we
		run our program
	*/
	fmt.Println("Hello World")
}
