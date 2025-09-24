package analysis

import (
	"analysis-question-AI/internal/core"
	"analysis-question-AI/internal/entity"
)

type analysisCommand struct {
	svc 	*analysisService
	log 	*core.Logger
	cfg 	*core.Config
}


func newAnalysisCommand(svc *analysisService, log *core.Logger, cfg *core.Config) *analysisCommand {
	return &analysisCommand{
		svc: svc,
		log: log,
		cfg: cfg,	
	}
}

func (a *analysisCommand) AnalyzeQuestions(questions []entity.QuestionWithRow) []entity.QuestionWithRow {

	analyzedQuestions, err := a.svc.findWrongQuestions(questions)
	if err != nil {
		a.log.Fatal(err)
		panic(err)
	}
	

	return analyzedQuestions
}