package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {
	// Read the input file
	inputFile, err := os.Open("./problems.csv")
	if err != nil {
		fmt.Println("Unable to read input file")
		os.Exit(1)
	}

	// Close the file while exiting this function
	defer inputFile.Close()

	// Result variables
	correctAnswers,totalQuestions := 0,0 

	// Read the csv records
	csvReader := csv.NewReader(inputFile)
	for {
		record, err := csvReader.Read()
		// Handle EOF error
		if err == io.EOF {
			break
		}
		// All other errors
		if err != nil {
			fmt.Println("Error file parsing input file: "+err.Error())
			inputFile.Close()
			os.Exit(1)
		}

		// Print the question
		fmt.Print(record[0]," ")
		totalQuestions++

		// Get the answer from user
		var answer string
		fmt.Scanln(&answer)

		// Check the answer
		if answer==record[1] {
			correctAnswers++
		}
	}

	fmt.Printf("%d out of %d questions correct\n",correctAnswers,totalQuestions)
}