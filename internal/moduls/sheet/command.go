package sheet

import (
	"analysis-question-AI/internal/api/cli"
	"analysis-question-AI/internal/core"
)

type sheetCommand struct {
	svc 	*sheetService
	flags 	*cli.flagsConfig
	log 	*core.Logger
}

func newSheetCommand(log *core.Logger, flags *cli.flagsConfig, svc *sheetService) *sheetCommand {
	return &sheetCommand{
		svc: svc,
		flags: flags,
		log: log,
	}
}


func (s *sheetCommand) AllQuestions() []QuestionWithRow {
	question, err := s.svc.getQuestions()
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	return question
}