package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

const fileName = "problems.csv"

func main() {
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

	var answer string
	var result int

	for i, quest := range questions {
		q, ans := quest[0], quest[1]
		fmt.Printf("\n%d. %s = ", i+1, q)
		if _, err := fmt.Scan(&answer); err != nil {
			fmt.Printf("failed to scan %v", err)

		}

		if answer == ans {
			result++
		}
	}

	fmt.Printf("your score: %d/%d\n", result, len(questions))
}
