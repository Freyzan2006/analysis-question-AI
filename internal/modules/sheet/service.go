package sheet

import (
	"analysis-question-AI/internal/core"
	"analysis-question-AI/internal/entity"
)

type sheetService struct {
	repo *sheetRepository
	log  *core.Logger
}

func newSheetService(repo *sheetRepository, log *core.Logger) *sheetService {
	return &sheetService{
		repo: repo,
		log:  log,
	}
}

func (s *sheetService) getQuestions() ([]entity.QuestionWithRow, error) {
	questions, err := s.repo.findAll()
	if err != nil {
		s.log.Fatal(err)
	}

	return questions, nil
}

func (s *sheetService) updateQuestions(questions []entity.QuestionWithRow) error {
	return s.repo.UpdateQuestionBlock(questions)
}
