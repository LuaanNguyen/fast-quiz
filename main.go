package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	//fmt.Println("Hello World!")

	file, err := os.Open("data/problems.csv")
	if err != nil {
		log.Fatal("Error while reading the file", err)
	}

	defer file.Close() //Ensure that the file is properly closed after the function is completed 

	reader := csv.NewReader(file) 

	problems, err := reader.ReadAll()

	if err != nil {
		fmt.Println("Error reading problems")
	}

	for _, problem := range problems {
		fmt.Println(problem)
	}
}