package sheet 


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

type QuestionWithRow struct {
    QuestionTable
    SheetName string // из какого листа
    StartRow  int    // 1-based: первая строка 4-строчного блока
}