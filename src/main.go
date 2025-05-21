package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"math/rand/v2"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("questions.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	reader := csv.NewReader(file)
	problems, err := reader.ReadAll()
	rand.Shuffle(len(problems), func(i, j int) {
		problems[i], problems[j] = problems[j], problems[i]
	})

	fmt.Println("Shuffled problems", problems)
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return
	}

	score := 0
	scanner := bufio.NewScanner(os.Stdin)

	for i, line := range problems {
		question := line[0]
		answer := strings.TrimSpace(line[1])

		fmt.Printf("Problem #%d: %s = ", i+1, question)
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		if input == answer {
			score++
		}
	}

	fmt.Printf("\nYou scored %d out of %d!\n", score, len(problems))
}
