package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
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
	user := map[string]string{
		"X": "Rock",
		"Y": "Paper",
		"Z": "Scissors",
	}

	// // points
	round := map[string]int{
		"win":  6,
		"draw": 3,
		"loss": 0,
	}
	shape := map[string]int{
		"Rock":     1,
		"Paper":    2,
		"Scissors": 3,
	}

	for scanner.Scan() {
		line := scanner.Text()
		moves := strings.Split(string(line), " ")
		opponentsMove := shape[opponent[moves[0]]]
		userMove := shape[user[moves[1]]]

		// first player (opponent) win would 1
		//  second player (user) would be 2
		// tie would be zero
		outcome := ((opponentsMove - userMove + 3) % 3)

		currentScore += userMove
		if outcome == 2 {
			currentScore += round["win"]
		} else if outcome == 1 {
			currentScore += round["loss"]
		} else if outcome == 0 {
			currentScore += round["draw"]
		}

		return currentScore
	}
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
}
