package entity

type QuestionWithRow struct {
	QuestionTable
	SheetName string
	StartRow  int
}