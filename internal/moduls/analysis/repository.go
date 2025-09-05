package analysis


type analysisRepository struct {
	api *analysisApi
}

func newAnalysisRepository(api *analysisApi) *analysisRepository {
	return &analysisRepository{
		api: api,
	}
}


func (a *analysisApi) analyzeQuestions() (*QuestionTable, bool, error) {
	analyzed, changed, err := a.api.analyzeQuestions() 
	if err != nil {
		return nil, false, err
	}

	return analyzed, changed, nil
}
