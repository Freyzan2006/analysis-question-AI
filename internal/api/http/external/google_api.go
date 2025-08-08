package external

import (
	"fmt"
)

import (
	"analysis-question-AI/internal/model"
)

type GoogleDocsAPI struct {
	URL string
}

func NewGoogleDocsAPI(url string) *GoogleDocsAPI {
	return &GoogleDocsAPI{
		URL: url,
	}
}


func(g *GoogleDocsAPI) GetQuestions() (*[]model.Question, error) {
	fmt.Println(g.URL)
	
	return nil, nil
}