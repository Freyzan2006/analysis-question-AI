package service 

import (
	"fmt"
	"log"
)

import (
	"analysis-question-AI/internal/model"
	"analysis-question-AI/internal/repository"
	"analysis-question-AI/internal/api/http/external"
)

import (
	"analysis-question-AI/internal/core"
)


type QuestionService struct {
	api  external.GeminiAPI
	repo repository.QuestionRepository
	svc  GoogleDocsService
	log  *core.Logger
}

func NewQuestionService(api *external.GeminiAPI, repo *repository.QuestionRepository, svc *GoogleDocsService, log *core.Logger) *QuestionService {
	return &QuestionService{
		repo: *repo,
		api:  *api,
		svc:  *svc,
		log:  log,
	}
}




func (s *QuestionService) Send() ([]model.QuestionTable, error) {
    questions, err := s.svc.GetQuestions()
    if err != nil {
        return nil, fmt.Errorf("ошибка получения вопросов: %w", err)
    }

	s.log.Info("Получено вопросов:", len(questions))	
	

	var results []model.QuestionTable

	for _, q := range questions {
        analyzed, changed, err := s.api.GenerateText(q.QuestionTable)
        if err != nil { return nil, err }

        if changed {
            row := q.StartRow 
            if err := s.svc.UpdateQuestionBlock(q.SheetName, row, *analyzed); err != nil {
                log.Printf("Ошибка обновления '%s'!A%d:E%d: %v", q.SheetName, row, row+3, err)
            }

			s.log.Info("Обновлен вопрос в листе ", q.SheetName, " Блок с строкой ", row, " - ", row+3)
        } 

        results = append(results, *analyzed)
    }




    if err := s.repo.Save(results, "./answers.md"); err != nil {
        return nil, fmt.Errorf("ошибка сохранения: %w", err)
    }

	if err := s.repo.SaveJSON(results, "./answers.json"); err != nil {
		return nil, fmt.Errorf("ошибка сохранения: %w", err)
	}



    return results, nil
}
