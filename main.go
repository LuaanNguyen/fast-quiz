package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

const DATA_PATH = "data/problems.csv"
const DEFAULT_TIME = 30

type problem struct {
	question string 
	answer string
}

func main() {
	// Open and read data from csv
	problems, err := readProblems(DATA_PATH)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize timer 
	timeLimit := time.Duration(DEFAULT_TIME) * time.Second
	timer := time.NewTimer(timeLimit)

	fmt.Printf("You have %d seconds to finish this quiz.\n", DEFAULT_TIME)
	fmt.Println("Press Enter to start...")
	fmt.Scanln() // Wait for the enter key 

	var correct, incorrect int 
	finished := make(chan bool)

	go func() {
		for i, problem := range problems {
			fmt.Printf("%d. Question: %s\n", i + 1, problem.question)
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
		finished <- true
	} ()

	select {
	case <- timer.C:
		fmt.Println("Time's up")
	case <- finished:
		fmt.Println("Quiz completed") 
	}
	
	fmt.Println("---------------------------------------------")
	fmt.Printf("Correct: %d, Incorrect: %d\n", correct, incorrect)
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
		fmt.Printf("Reading from file: %s\n", fileName)
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