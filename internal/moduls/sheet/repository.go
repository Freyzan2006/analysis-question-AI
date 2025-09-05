package sheet 

type sheetRepository struct {
	api *sheetApi
}

func newSheetRepository() *sheetRepository {
	return &sheetRepository{}
}

func (s *sheetRepository) findAll() ([]QuestionWithRow, error) {
	return s.api.getQuestions()
}