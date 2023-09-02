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

func day1(fileName string) int {
	file, err := os.Open(fileName)
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
	for _, calory := range allCalories {
		if calory > max {
			max = calory
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return max
}

func main() {
	maxCalory := day1("day1.txt")
	fmt.Println("maxCalory: ", maxCalory)
}
