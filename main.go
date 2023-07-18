package main

import (
	"fmt" 
	"os"
)

func main() {
	// basic stuff
	fmt.Println("Your lack of faith disturbs me")

	
	// ! LOOPS
	// i, arg := range os.Args[1:]
	// similar to for index, item in
	for i, arg := range os.Args[1:] {
		fmt.Println(i, arg)
	}

	// initialization; condition; increment
	// The firs t element of os.Args, os.Args[0], is the name of the command its elf; 
	for i := 0; i < len(os.Args); i++ {
		fmt.Println(os.Args[i])
	}
}