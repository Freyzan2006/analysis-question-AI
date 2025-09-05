package sheet 

type sheetService struct {
	repo *sheetRepository
}

func newSheetService() *sheetService {
	return &sheetService{}
}


func (s *sheetService) getQuestions() ([]QuestionWithRow, error) {
	questions, err := s.repo.findAll()
	if err != nil {
		log.Fatal(err)
	}

	return questions, nil
}