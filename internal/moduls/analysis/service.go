package analysis


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

func findWrongQuestions(questions []QuestionTable) (QuestionTable, error) {

	var results []model.QuestionTable

	for _, q := range questions {
        analyzed, changed, err := a.repo.analyzeQuestions(q.QuestionTable)
        if err != nil { return nil, err }

        if changed {
            // row := q.StartRow 
            // if err := a.svc.CorrectQuestions(questions); err != nil {
            //     a.log.Printf("Ошибка обновления '%s'!A%d:E%d: %v", q.SheetName, row, row+3, err)
            // }

			// a.log.Info("Обновлен вопрос в листе ", q.SheetName, " Блок с строкой ", row, " - ", row+3)

			results = append(results, *analyzed)
			a.log("Не корректный вопрос в листе:", q.SheetName, " Строка: начиная", q.StartRow, "По концу", q.StartRow+3)
        } 
    }

	a.log("Найдено неправильных вопросов:", len(results))

	return results, nil
}


