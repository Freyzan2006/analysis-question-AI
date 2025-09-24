package main


import (
	"analysis-question-AI/internal/app"
)



func main() {	
	application := app.NewCheckingCorrectingQuestionApplication()
	application.Run()
}
