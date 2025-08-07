package repository

import (
	"fmt"
)

import (
	"analysis-question-AI/internal/model"
)

type QuestionRepository struct {}

func NewQuestionRepository() *QuestionRepository {
	return &QuestionRepository{}
}


func (r *QuestionRepository) Save(question model.Question, formatSave string) (*model.Question, error) {
	fmt.Println("Всё успешно сохранено в", formatSave)

	return &model.Question{
		Question: question.Question,
		Answer: question.Answer,
	}, nil
}
