package service 

import (
	"log"
	"fmt"
)

import (
	"analysis-question-AI/internal/model"
	"analysis-question-AI/internal/repository"
	"analysis-question-AI/internal/api/http/external"
)


type QuestionService struct {
	api  external.GeminiAPI
	repo repository.QuestionRepository
	svc  GoogleDocsService
}

func NewQuestionService(api *external.GeminiAPI, repo *repository.QuestionRepository, svc *GoogleDocsService) *QuestionService {
	return &QuestionService{
		repo: *repo,
		api:  *api,
		svc:  *svc,
	}
}




func (s *QuestionService) Send() ([]model.QuestionTable, error) {
    questions, err := s.svc.GetQuestions()
    if err != nil {
        return nil, fmt.Errorf("ошибка получения вопросов: %w", err)
    }

	log.Println("Какие вопросы получены:", questions)


	var results []model.QuestionTable
	// for _, q := range questions {
	// 	analyzed, changed, err := s.api.GenerateText(q)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	if changed {
	// 		log.Printf("Вопрос обновлён: %s → %s\n", q.Question, analyzed.Question)
	// 	} else {
	// 		log.Printf("Вопрос без изменений: %s\n", q.Question)
	// 	}

	// 	results = append(results, *analyzed)
	// }
	for _, q := range questions {
        analyzed, changed, err := s.api.GenerateText(q.QuestionTable)
        if err != nil { return nil, err }

        if changed {
            row := q.StartRow // уже 1-based
            if err := s.svc.UpdateQuestionBlock(q.SheetName, row, *analyzed); err != nil {
                log.Printf("Ошибка обновления '%s'!A%d:E%d: %v", q.SheetName, row, row+3, err)
            }
        }

        results = append(results, *analyzed)
    }




    if err := s.repo.Save(results, "./answers.md"); err != nil {
        return nil, fmt.Errorf("ошибка сохранения: %w", err)
    }

	if err := s.repo.SaveJSON(results, "./answers.json"); err != nil {
		return nil, fmt.Errorf("ошибка сохранения: %w", err)
	}

	// if err := s.repo.SaveToSheets(results, "Answers"); err != nil {
	// 	return nil, fmt.Errorf("ошибка сохранения в Google Sheets: %w", err)
	// }


    return results, nil
}
