package Play

import (
	"flag"
	"fmt"
	"time"
)

var changeStatus = make(chan int)

func StartQuiz() {
	fmt.Println("\tWELCOME TO THE QUIZ GAME")
	fmt.Println("LET's BEGIN!")
	csvFileName := flag.String("csv", "./problems.csv", "a csv file in the format of `question, answer`")
	timeLimit := flag.Int("limit", 30, "the time limit for quiz in seconds")
	flag.Parse()

	records := readFromCSV(csvFileName)
	problems := getQuestionsAndAnswers(records)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	getUservalidations(problems, timer)
}
