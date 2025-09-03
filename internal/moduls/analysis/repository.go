package analysis


type analysisRepository struct {
	api *analysisApi
}

func newAnalysisRepository (api *analysisApi) *analysisRepository {
	return &analysisRepository{
		api: api,
	}
}