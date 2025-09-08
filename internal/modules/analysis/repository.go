package analysis

import (
	"analysis-question-AI/internal/entity"
)


type analysisRepository struct {
	api *analysisApi
}

func newAnalysisRepository(api *analysisApi) *analysisRepository {
	return &analysisRepository{
		api: api,
	}
}


func (a *analysisRepository) analyzeQuestions(q entity.QuestionTable) (*entity.QuestionTable, bool, error) {
	analyzed, changed, err := a.api.analyzeQuestionsApi(q) 
	if err != nil {
		return nil, false, err
	}

	return analyzed, changed, nil
}
