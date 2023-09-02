package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("day1.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	allCalories := []int{}
	totalCalories := 0

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			allCalories = append(allCalories, totalCalories)
			totalCalories = 0
			continue
		}
		num, err := strconv.Atoi(line)
		check(err)
		totalCalories += num
	}
	max := allCalories[0]
	for index, calory := range allCalories {
		fmt.Println("calory: ", calory)
		if calory > max {
			max = calory
		}
		fmt.Println("index: ", index)
	}

	fmt.Println(max)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
