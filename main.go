package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Question struct {
	question string
	answer   string
}

func main() {

	var score int = 0

	questions := getFileDetials()

	for i := 0; i < len(questions); i++ {

		fmt.Print("\n", questions[i].question)

		ans := string(questions[i].answer)

		x := getInputDetails()

		if strings.Compare(ans, x) == 0 {
			score = score + 1
			fmt.Println("correct ✅")
		} else {
			fmt.Println("wrong 👎")
		}

		continue

	}

	fmt.Println("this is your final score: ", score, "/", len(questions))

}

func getFileDetials() []Question {

	fileName := "questions.txt"

	data, err := ioutil.ReadFile(fileName)

	var y []string

	if err != nil {

		fmt.Println("an error occured ", err, " 😵‍💫")

	} else {
		x := string(data)
		y = strings.Split(x, "\n")
	}

	allQuestions := []Question{}

	for i := 0; i < len(y); i++ {
		var d = strings.Split(y[i], ",")
		t := Question{
			question: d[0],
			answer:   strings.ReplaceAll(d[1], " ", ""),
		}
		allQuestions = append(allQuestions, t)
	}

	return allQuestions

}

func getInputDetails() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\nEnter text: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSuffix(text, "\n")
	return text
}
