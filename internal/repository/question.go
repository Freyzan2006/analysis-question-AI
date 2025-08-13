// package repository

// import (
// 	"fmt"
// )

// import (
// 	"analysis-question-AI/internal/model"
// )

// type QuestionRepository struct {}

// func NewQuestionRepository() *QuestionRepository {
// 	return &QuestionRepository{}
// }


// func (r *QuestionRepository) Save(question []model.QuestionTable, formatSave string) error {
// 	fmt.Println("Всё успешно сохранено в", formatSave)

	

// 	return nil
// }

package repository

import (
	"fmt"
	"os"
	"strings"

	"analysis-question-AI/internal/model"
)

type QuestionRepository struct{}

func NewQuestionRepository() *QuestionRepository {
	return &QuestionRepository{}
}

func (r *QuestionRepository) Save(questions []model.QuestionTable, formatSave string) error {
	var sb strings.Builder

	for i, q := range questions {
		sb.WriteString(fmt.Sprintf("### Вопрос %d\n", i+1))
		sb.WriteString(fmt.Sprintf("%s\n\n", q.Question))

		for j, opt := range q.Options {
			sb.WriteString(fmt.Sprintf("%d) %s\n", j+1, opt.Text))
			if opt.IsCorrect {
				sb.WriteString("✅ Правильный ответ\n")
			} else {
				sb.WriteString("❌ Неправильный ответ\n")
			}
			if opt.Explanation != "" {
				sb.WriteString(fmt.Sprintf("_Пояснение_: %s\n", opt.Explanation))
			}
			sb.WriteString("\n")
		}
		sb.WriteString("\n---\n\n")
	}

	// Записываем в файл
	if err := os.WriteFile(formatSave, []byte(sb.String()), 0644); err != nil {
		return fmt.Errorf("ошибка записи в файл: %w", err)
	}

	fmt.Println("Всё успешно сохранено в", formatSave)
	return nil
}

