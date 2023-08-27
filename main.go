package main

import "fmt"

func main() {
	fmt.Println("Hello, Gophers!")
	//datatypes()
	//arithmetic()
	//comparison()
	//pointers()
	cli()
}

func datatypes() {
	var a string
	a = "Akshay \n" + `b \n`
	fmt.Print(a)

	var b int = 99
	fmt.Println(b)

	var c bool = true
	fmt.Println(c)

	d := 3.14
	fmt.Println(d)

	var e int = int(d)
	fmt.Println(e)
}

func arithmetic() {
	a, b := 10, 5

	c := a + b
	c = a - b
	c = a * b
	c = a / b
	c = a % 3
	fmt.Println(c)

	d := 10.5 / .5

	fmt.Println(d)
}

func comparison() {

	a, b := 10, 5

	c := a == b
	c = a != b
	c = a < b
	c = a <= b
	c = a > b
	c = a >= b
	fmt.Print(c)
}

func pointers() {
	a := "foo" // create a string variable

	b := &a // address operator returns address of a variable

	fmt.Println(*b) // dereferenced a pointer using asterik
	*b = "bar"

	fmt.Println(a)
}
