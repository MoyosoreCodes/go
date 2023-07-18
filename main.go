package main

import (
	"fmt" 
)

func main() {
	// basic stuff
	fmt.Println("Your lack of faith disturbs me")

	
	// ! LOOPS
	// i, arg := range os.Args[1:]
	// similar to for index, item in
	// for i, arg := range os.Args[1:] {
	// 	fmt.Println(i, arg)
	// }

	// initialization; condition; increment
	// The first element of os.Args, os.Args[0], is the name of the command its elf; 
	// for i := 0; i < len(os.Args); i++ {
	// 	fmt.Println(os.Args[i])
	// }

	// fmt.Println(strings.Join(os.Args[1:], " "))
	// fmt.Println(os.Args[1:][0])

	// maps
	hashMap := make(map[string]string)
	hashMap["foo"] = "bar"
	fmt.Println(hashMap)

	for key, value := range hashMap {
		fmt.Println("key: ", key)
		fmt.Println("value: ", value)
	}
}