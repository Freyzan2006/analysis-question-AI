package model 


// type Answer struct {
// 	Question string `json:"question"`
// 	Answer string `json:"answer"`
// 	BeforeAnswer string `json:"beforeAnswer"`
// }




type AnswerOption struct {
	Text        string `json:"text"`
	IsCorrect   bool   `json:"isCorrect"`
	Explanation string `json:"explanation"`
}

type QuestionTable struct {
	Question   string         `json:"question"`
	Options    []AnswerOption `json:"options"`
	Categories []string       `json:"categories"` // <-- добавили
}
