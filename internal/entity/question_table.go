package entity

type QuestionTable struct {
	Question   string         `json:"question"`
	Options    []AnswerOption `json:"options"`
	Categories []string       `json:"categories"` // <-- добавили
}