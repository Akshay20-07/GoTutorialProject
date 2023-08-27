package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cli() {

	fmt.Println("what would do you like to have? \n 1. Coffee \n 2. Beer \n 3. Tea")
	in := bufio.NewReader(os.Stdin)
	s, _ := in.ReadString('\n')
	s = strings.TrimSpace(s)
	s = strings.ToUpper(s)

	fmt.Println(s)

}
