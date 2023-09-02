package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"fmt"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getCaloriesTotal(fileName string) []int {
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

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return allCalories
}

func getTopCalories(array []int, top int) []int {
	if len(array) < top {
		return array
	}

	sort.Sort(sort.Reverse(sort.IntSlice(array)))

	return array[:top]
}

func main() {
	calories := getCaloriesTotal("day1.txt")
	topThree := getTopCalories(calories, 3)
	fmt.Println("max: ", topThree[0])
	fmt.Println("topThree: ", topThree)
	sum := 0
	for _, calory := range topThree {
		sum += calory
	}
	fmt.Println("topThree total: ", sum)
}
