package service 

import (
	"log"
)

import (
	"analysis-question-AI/internal/model"
	"analysis-question-AI/internal/repository"
	"analysis-question-AI/internal/api/http/external"
)


type QuestionService struct {
	api  external.GeminiAPI
	repo repository.QuestionRepository
}

func NewQuestionService(api *external.GeminiAPI, repo *repository.QuestionRepository) *QuestionService {
	return &QuestionService{
		repo: *repo,
		api:  *api,
	}
}

func (s *QuestionService) Send(question model.Question) (*model.Question, error) {
	answer, err := s.api.GenerateText(question.Question)
	if err != nil {
		log.Fatal(err)
	}

	saved, err := s.repo.Save(*answer, "./answers.json")
	if err != nil {
		log.Fatal(saved)
	}

	return answer, nil
}