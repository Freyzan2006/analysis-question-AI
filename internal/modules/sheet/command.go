package sheet

import (
	"analysis-question-AI/internal/core"
	"analysis-question-AI/internal/entity"
)

type sheetCommand struct {
	svc 	*sheetService
	cfg 	*core.Config
	log 	*core.Logger
}

func newSheetCommand(log *core.Logger, cfg *core.Config, svc *sheetService) *sheetCommand {
	return &sheetCommand{
		svc: svc,
		cfg: cfg,
		log: log,
	}
}


func (s *sheetCommand) AllQuestions() []entity.QuestionWithRow {
	question, err := s.svc.getQuestions()
	if err != nil {
		s.log.Fatal(err)
		panic(err)
	}

	return question
}


func (s *sheetCommand) SaveQuestions(questions []entity.QuestionWithRow) {
	err := s.svc.updateQuestions(questions)
	if err != nil {
		s.log.Fatal(err)
		panic(err)
	}

}