package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

type Problem struct {
	q string
	a string
}

func quizTheUser(problem Problem) bool {
	// Print the question
	fmt.Print(problem.q, " ")

	// Get the answer from user
	var answer string
	fmt.Scanf("%s\n", &answer)

	// Validate the answer
	if answer == problem.a {
		return true
	} else {
		return false
	}
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func main() {

	// Define input file flag
	inputFileName := flag.String("file", "problems.csv", "Name of the input file. File should be present in the current folder.")

	// Parse the flags
	flag.Parse()

	// Read the input file
	inputFile, err := os.Open(*inputFileName)
	if err != nil {
		exit("Unable to read input file")
	}

	// Close the file while exiting this function
	defer inputFile.Close()

	// Result variables
	correctAnswers, totalQuestions := 0, 0

	// Read the csv records one by one
	csvReader := csv.NewReader(inputFile)
	for {
		record, err := csvReader.Read()
		// Handle EOF error
		if err == io.EOF {
			break
		}
		// All other errors
		if err != nil {
			inputFile.Close()
			exit("Error file parsing input file: " + err.Error())
		}

		totalQuestions++

		// Convert the input to Problem type
		newProblem := Problem{
			q: record[0],
			a: strings.TrimSpace(record[1]),
		}

		// Check the answer
		if quizTheUser(newProblem) {
			correctAnswers++
		}
	}

	fmt.Printf("%d out of %d questions correct\n", correctAnswers, totalQuestions)
}
