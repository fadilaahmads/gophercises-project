package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func timer(timeout int, ch chan<- bool) {
	time.AfterFunc(time.Duration(timeout)*time.Second, func() {
		ch <- true
	})
}

func watcher(timeout int, ch <-chan bool) {
	<-ch
	fmt.Println("Timeout! no answer after", timeout, " seconds")
	os.Exit(0)
}

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

	// timer
	var timeout = 5
	var ch = make(chan bool)

	correct_answer := 0
	for _, value := range data {
		question := value[0]
		answer := value[1]

		go timer(timeout, ch)
		go watcher(timeout, ch)

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
