package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
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

	var correct int = 0 
	var incorrect int = 0 

	for i, problem := range problems {
		log.Printf("%d. Question: %s\n", i + 1, problem.question)
		var input string
		fmt.Scanln( &input)
		if strings.TrimSpace(input) == strings.TrimSpace(problem.answer) {
			correct ++
			fmt.Println("Correct!")
		} else {
			incorrect ++
			fmt.Println("Incorrect!")
		}
		fmt.Println()
	}
	
	log.Printf("Correct: %d, Incorrect: %d", correct, incorrect)
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
	} else {
		log.Printf("Reading from file: %s", fileName)
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