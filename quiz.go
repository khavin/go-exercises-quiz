package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
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

func convertCSVToProblems(records [][]string) []Problem {
	problems := make([]Problem, len(records))

	for i := range records {
		problems[i] = Problem{
			q: records[i][0],
			a: strings.TrimSpace(records[i][1]),
		}
	}

	return problems
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func main() {

	// Define the flags
	inputFileName := flag.String("file", "problems.csv", "Name of the input file. File should be present in the current folder.")
	timeout := flag.Int("timeout", 30, "Quiz timeout period.")

	// Parse the flags
	flag.Parse()

	// Read the input file
	inputFile, err := os.Open(*inputFileName)
	if err != nil {
		exit("Unable to read input file")
	}

	// Close the file while exiting this function
	defer inputFile.Close()

	// Read the csv records all at once
	csvReader := csv.NewReader(inputFile)
	records, err := csvReader.ReadAll()

	if err != nil {
		inputFile.Close()
		exit("Error file parsing input file: " + err.Error())
	}

	// Result variables
	correctAnswers, totalQuestions := 0, len(records)

	// Convert the input records to Problem type
	problems := convertCSVToProblems(records)

	// Set a timer
	timer := time.NewTimer(time.Duration(*timeout) * time.Second)
	timedOut := false

	for i := range problems {

		// This channel will be used to receive the problem result
		userChannel := make(chan bool)

		// Quiz the user
		go func() {
			userChannel <- quizTheUser(problems[i])
		}()

		// Wait until either the timer ends or the user responds
		select {
		case result := <-userChannel:
			if result {
				correctAnswers++
			}
		case <-timer.C:
			fmt.Println()
			timedOut = true
		}

		if timedOut {
			break
		}
	}

	fmt.Printf("%d out of %d questions correct\n", correctAnswers, totalQuestions)
}
