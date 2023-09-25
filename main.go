package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"math"
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

func getTotalGameScore(fileName string) int {
	file, err := os.Open(fileName)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// scores := []int{}
	currentScore := 0

	// // player moves
	opponent := map[string]string{
		"A": "Rock",
		"B": "Paper",
		"C": "Scissors",
	}
	user := map[string]map[string]string{
		"X": {"Rock": "C", "Paper": "A", "Scissors": "B"},
		"Y": {"Rock": "A", "Paper": "B", "Scissors": "C"},
		"Z": {"Rock": "B", "Paper": "C", "Scissors": "A"},
	}

	// // points
	outcome := map[string]int{
		"Z": 6,
		"Y": 3,
		"X": 0,
	}
	shape := map[string]int{
		"Rock":     1,
		"Paper":    2,
		"Scissors": 3,
	}

	for scanner.Scan() {
		line := scanner.Text()
		moves := strings.Split(string(line), " ")
		opponentsMove := opponent[moves[0]]
		roundOutcome := outcome[moves[1]]
		userMove := opponent[user[moves[1]][opponentsMove]]

		// first player (opponent) win would 1
		//  second player (user) would be 2
		// tie would be zero
		// outcome := ((opponentsMove - userMove + 3) % 3)

		currentScore += shape[userMove] + roundOutcome
		// if outcome == 2 {
		// 	currentScore += round["win"]
		// } else if outcome == 1 {
		// 	currentScore += round["loss"]
		// } else if outcome == 0 {
		// 	currentScore += round["draw"]
		// }

	}
	return currentScore
}

// func countingSort(arr []int32) []int32 {
//     // Write your code here
//     const length := len(arr)
//     frequency := [length]int32{}

//     for _, num := range(arr) {

//     }

//     return frequency

// }

func bubbleSort(arr []int32) []int32 {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				tmp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tmp
			}
		}
	}
	return arr
}

func binarySearch(arr []int, searchValue int) int {
	// assuming the arr is sorted
	low := 0
	high := len(arr)

	for low <= high {
		mid := low + (high-low)/2
		value := arr[mid]

		if value == searchValue {
			return mid
		} else if value > searchValue {
			high = mid
		} else {
			low = mid + 1
		}
	}

	return -1
}

func crystalBall(breaks []bool) float64 {
	// Question: Given two crystal balls that will  break if dropped from high enough
	// distance, determine the exact spot in which it will break in the most 
	// optimized way

	// params breaks: boolean[]
	// returns the index
	jmpAmount := math.Floor(math.Sqrt(float64(len(breaks))))
	return jmpAmount
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

	scores := getTotalGameScore("day2.txt")
	fmt.Println(scores)

	// valuesToSort := []int32{5, 4, 3, 2, 1}
	fmt.Println(bubbleSort([]int32{4, 5, 2, 1, 3}))
	fmt.Println(binarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, 11))
	fmt.Println(crystalBall([]bool{false, false, true}))
}
