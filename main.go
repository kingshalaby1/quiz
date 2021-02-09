package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

const fileName = "problems.csv"
var result *int

func main() {

	result = new(int)
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("failed to open file: %v\n", err)
		return
	}
	defer f.Close()

	r := csv.NewReader(f)
	questions, err := r.ReadAll()
	if err != nil {
		fmt.Printf("failed to open CSV file: %v\n", err)
		return
	}

	finished := startQuiz(questions)
	timeUp := time.Tick(12 * time.Second)

	select{
	case <-finished:
	case <-timeUp:
	}

	fmt.Printf("your score: %d/%d\n", *result, len(questions))
}

func startQuiz(questions [][]string) chan bool{
	var answer string

	finished := make(chan bool)

	go func() {
		for i, quest := range questions {
			q, ans := quest[0], quest[1]
			fmt.Printf("\n%d. %s = ", i+1, q)
			if _, err := fmt.Scan(&answer); err != nil {
				fmt.Printf("failed to scan %v", err)

			}

			if answer == ans {
				*result++
			}
		}
		finished <- true
	}()
	return finished
}
