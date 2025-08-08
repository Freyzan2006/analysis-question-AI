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

type googleDocsService struct {
	api  external.GoogleDocsAPI
	repo repository.QuestionRepository
}

func NewGoogleDocsService(api *external.GoogleDocsAPI, repo *repository.QuestionRepository) *googleDocsService {
	return &googleDocsService{
		repo: *repo,
		api:  *api,
	}
}

func (g *googleDocsService) GetQuestions() (*model.Question, error) {
	
	api, err := g.api.GetQuestions()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(api)
	
	return nil, nil
}