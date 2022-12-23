package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	// open CSV file
	csvfile := flag.String("csvfile", "problems.csv", "Quiz game content")
	flag.Parse()

	file, err := os.Open(*csvfile)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()
	csvReader := csv.NewReader(file)

	data, err := csvReader.ReadAll()
	correct_answer := 0
	for _, value := range data {
		question := value[0]
		answer := value[1]

		fmt.Println("Question: ", question)
		var input string
		fmt.Scanln(&input)
		// fmt.Println("Your answer: ", input)
		// Checking input
		if input != answer {
			fmt.Println("Wrong Answer!")
		} else {
			fmt.Println("Correct!")
			correct_answer++
		}
	}
	fmt.Printf("Score: %d/%d", correct_answer, len(data))

}
