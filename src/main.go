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

	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return
	}

	score := 0
	scanner := bufio.NewScanner(os.Stdin)

	for i, line := range problems {
		attemptsLeft := 3
		for attemptsLeft > 0 {
			question := line[0]
			answer := strings.TrimSpace(line[1])

			fmt.Printf("Problem #%d: %s = ", i+1, question)
			scanner.Scan()
			input := strings.TrimSpace(scanner.Text())

			if input == answer {
				score++
				break
			} else {
				attemptsLeft--
				if attemptsLeft == 0 {
					fmt.Printf("You ran out of tries for this question. The correct answer is %s.\n", answer)
					break
				} else {
					fmt.Printf("Wrong answer. You have %d tries left.\n", attemptsLeft)
				}
			}
		}
	}

	fmt.Printf("\nYou scored %d out of %d!\n", score, len(problems))
}
