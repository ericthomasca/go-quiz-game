package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	csvFileName := flag.String("csv", "problems.csv", "A string with the filename of the CSV including extension")
	flag.Parse()

	file, err := os.Open(*csvFileName)

	if err != nil {
		fmt.Printf("Could not find file named %s\n", *csvFileName)
		os.Exit(1)
	}

	csv_data := csv.NewReader(file)

	problems, err := csv_data.ReadAll()

	if err != nil {
		fmt.Println("Problem with CSV data")
		os.Exit(1)
	}

	for _, problem := range problems {
		fmt.Println(problem)
	}

}
