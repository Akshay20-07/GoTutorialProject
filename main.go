package main

import "fmt"

func main() {
	fmt.Println("Hello, Gophers!")
	datatypes()
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
