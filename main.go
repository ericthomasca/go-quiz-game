package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

type problem struct {
	question string
	answer string
}

func parseLines(lines [][]string) []problem {
	problemChan := make([]problem, len(lines))
	for i, line := range lines {
		problemChan[i] = problem{
			question: line[0],
			answer: line[1],
		}
	}
	return problemChan
}

func main() {
	csvFileName := flag.String("csv", "problems.csv", "A string with the filename of the CSV including extension")
	flag.Parse()

	file, err := os.Open(*csvFileName)
	if err != nil {
		fmt.Printf("Could not find file named %s\n", *csvFileName)
		os.Exit(1)
	}

	csv_data := csv.NewReader(file)

	lines, err := csv_data.ReadAll()
	if err != nil {
		fmt.Println("Problem with CSV data")
		os.Exit(1)
	}

	problems := parseLines(lines)

	for i, problem := range problems {
		fmt.Printf("Question %d: %s = %s\n", i+1, problem.question, problem.answer)
	}

}
