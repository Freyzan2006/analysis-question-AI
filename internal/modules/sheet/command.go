package sheet

import (
	"analysis-question-AI/internal/api/cli"
	"analysis-question-AI/internal/core"
	"analysis-question-AI/internal/entity"
)

type sheetCommand struct {
	svc 	*sheetService
	flags 	*cli.FlagsConfig
	log 	*core.Logger
}

func newSheetCommand(log *core.Logger, flags *cli.FlagsConfig, svc *sheetService) *sheetCommand {
	return &sheetCommand{
		svc: svc,
		flags: flags,
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