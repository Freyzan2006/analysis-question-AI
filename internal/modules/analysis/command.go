package analysis

import (
	"analysis-question-AI/internal/api/cli"
	"analysis-question-AI/internal/core"
	"analysis-question-AI/internal/entity"
)

type analysisCommand struct {
	svc 	*analysisService
	log 	*core.Logger
	flags 	*cli.FlagsConfig

}


func newAnalysisCommand(svc *analysisService, log *core.Logger, flags *cli.FlagsConfig) *analysisCommand {
	return &analysisCommand{
		svc: svc,
		log: log,
		flags: flags,	
	}
}

func (a *analysisCommand) AnalyzeQuestions(questions []entity.QuestionWithRow) []entity.QuestionTable {

	analyzedQuestions, err := a.svc.findWrongQuestions(questions)
	if err != nil {
		a.log.Fatal(err)
		panic(err)
	}
	

	return analyzedQuestions
}