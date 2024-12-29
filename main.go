package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type problem struct {
	question string 
	answer string
}

func main() {
	// Open and read data from csv
	problems, err := readProblems("data/problems.csv")
	if err != nil {
		log.Fatal(err)
	}

	for _, problem := range problems {
		fmt.Printf("Question: %s, Answer: %s\n", problem.question, problem.answer)
	}
}

func readProblems(fileName string) ([]problem, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close() //Ensure that the file is properly closed after the function is completed 

	reader := csv.NewReader(file) 
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	problems := make([]problem, len(records))
	for i, record := range records {
		problems[i] = problem{
			question: record[0],
			answer: record[1],
		}
	}

	return problems, nil 
}