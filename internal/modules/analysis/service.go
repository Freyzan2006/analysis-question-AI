package analysis

import (
	"analysis-question-AI/internal/core"
	"analysis-question-AI/internal/entity"
)

type analysisService struct {
	repo *analysisRepository
	log  *core.Logger
}

func newAnalysisService(repo *analysisRepository, log *core.Logger) *analysisService {
	return &analysisService{
		repo: repo,
		log:  log,
	}
}

func (a *analysisService) findWrongQuestions(questions []entity.QuestionWithRow) ([]entity.QuestionWithRow, error) {

	var results []entity.QuestionWithRow

	for _, q := range questions {
        analyzed, changed, err := a.repo.analyzeQuestions(q)
        if err != nil { return nil, err }

        if changed {
            // row := q.StartRow 
            // if err := a.svc.CorrectQuestions(questions); err != nil {
            //     a.log.Printf("Ошибка обновления '%s'!A%d:E%d: %v", q.SheetName, row, row+3, err)
            // }

			// a.log.Info("Обновлен вопрос в листе ", q.SheetName, " Блок с строкой ", row, " - ", row+3)

			results = append(results, *analyzed)
			a.log.Info("Не корректный вопрос в листе: ", q.SheetName, " Строка: начиная ", q.StartRow, "По концу ", q.StartRow+3)
        } 
    }

	a.log.Info("Найдено неправильных вопросов: ", len(results))

	return results, nil
}


