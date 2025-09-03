package analysis


type analysisService struct {
	repo *analysisRepository
}

func NewAnalysisService(repo *analysisRepository) *analysisService {
	return &analysisService{
		repo: repo,
	}
}