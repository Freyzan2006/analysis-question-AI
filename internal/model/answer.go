package model 


// type Answer struct {
// 	Question string `json:"question"`
// 	Answer string `json:"answer"`
// 	BeforeAnswer string `json:"beforeAnswer"`
// }




type AnswerOption struct {
	Text      string
	IsCorrect bool
	Explanation string
}

type QuestionTable struct {
	Question string
	Options  []AnswerOption
}