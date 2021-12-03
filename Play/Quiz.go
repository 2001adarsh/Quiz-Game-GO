package Play

import (
	"fmt"
)

func StartQuiz() {
	fmt.Println("\tWELCOME TO THE QUIZ GAME")
	fmt.Println("LET's BEGIN!")
	records := readFromCSV()
	questions, answers, total := getQuestionsAndAnswers(records)
	correct := getUservalidations(questions, answers)
	fmt.Printf("Your Result: %v out of %v", correct, total)
	return
}

func getUservalidations(questions []string, answers []int) int {
	correct := 0
	var input int
	for i:=0; i<len(questions); i++{
		fmt.Printf("Problem #%d: %s = ", i, questions[i])
		fmt.Scanln(&input)
		if input == answers[i] {
			correct++
		}
	}
	return correct
}

func getQuestionsAndAnswers(records [][]string) ([]string, []int, int) {
	total := len(records)
	var questions []string
	var answers []int

	for _, line := range records {
		questions = append(questions, line[0])
		answers = append(answers, convertToInt(line[1]))
	}
	return questions, answers, total
}
