package main

import (
	"fmt"
	"os"
)

func command_line_args() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func command_line_args_2() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

// Modif y the echo program to also print os.Args[0], the name of the command that invo ked it.
func command_line_args_E1() {
	s, sep := "", ""
	for _, arg := range os.Args[0:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

// Modif y the echo program to print the index and value of each of its arguments, one per line.
func command_line_args_E2() {
	s, sep := "", ""
	for i, arg := range os.Args[0:] {
		s += sep + arg
		sep = ""
		fmt.Printf("# %d: %s \n", i, s)
		s = ""
	}
}

func main() {
	fmt.Println("Hello, World!")
	command_line_args()
	// a "traditional" While loop
	m := 0
	for m <= 9 {
		fmt.Println("# ", m)
		m++
	}
	// A traditional infinite loop
	// for {
	//..
	//}
	// Command line args 2
	command_line_args_2()
	fmt.Println("Exercises.")
	command_line_args_E1()
	command_line_args_E2()
}
