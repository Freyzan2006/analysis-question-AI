package entity 

type AnswerOption struct {
	Text        string `json:"text"`
	IsCorrect   bool   `json:"isCorrect"`
	Explanation string `json:"explanation"`
}