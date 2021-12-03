package Play

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func convertToInt(a string) int {
	a = strings.TrimSpace(a)
	ans, err := strconv.Atoi(a)
	if err != nil {
		return -1
	}
	return ans
}

func startTimer(seconds *int) {
	for i := 0; i < *seconds; i++ {
		time.Sleep(time.Second)
	}
	changeStatus <- -1
}

func getUservalidations(problems []problem, timer *time.Timer) {
	correct, total := 0, len(problems)
	for i := 0; i < len(problems); i++ {
		fmt.Printf("Problem #%d: %s = ", i+1, problems[i].question)
		answerCh := make(chan int)
		go getInput(answerCh)
		select {
		case <-timer.C:
			fmt.Printf("\nYour Result: %v out of %v", correct, total)
			return
		case input := <-answerCh:
			if input == problems[i].answer {
				correct++
			}
		}
	}
	fmt.Printf("\nYour Result: %v out of %v", correct, total)
}

func getInput(answerChan chan int) {
	var input int
	fmt.Scanln(&input)
	answerChan <- input
}

func getQuestionsAndAnswers(records [][]string) []problem {
	var problems []problem

	for _, line := range records {
		question := line[0]
		answer := convertToInt(line[1])
		problems = append(problems, problem{question: question, answer: answer})
	}
	return problems
}
