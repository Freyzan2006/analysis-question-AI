package service 

import (
	"log"
)

import (
	"analysis-question-AI/internal/model"
	"analysis-question-AI/internal/repository"
	"analysis-question-AI/internal/api/http/external"
)

type GoogleDocsService struct {
	api  external.GoogleDocsAPI
	repo repository.QuestionRepository
}

func NewGoogleDocsService(api *external.GoogleDocsAPI, repo *repository.QuestionRepository) *GoogleDocsService {
	return &GoogleDocsService{
		repo: *repo,
		api:  *api,
	}
}

func (g *GoogleDocsService) GetQuestions() ([]model.QuestionTable, error) {
	
	api, err := g.api.GetQuestions()
	if err != nil {
		log.Fatal(err)
	}

	
	
	return api, nil
}