package analysis

import (
	"analysis-question-AI/internal/api/cli"
)

type analysisCommand struct {
	svc 	*analysisService
	log 	*core.Logger
	flags 	*cli.flagsConfig

}


func newAnalysisCommand(svc *analysisService, log *core.Logger, flags *cli.flagsConfig) *analysisCommand {
	return &analysisCommand{
		svc: svc,
		log: log,
		flags: flags,	
	}
}

func (a *analysisCommand) AnalyzeQuestions(questions []QuestionTable) []QuestionTable {

	questions, err := a.svc.findWrongQuestions(questions)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	

	return questions
}