package analysis

import (
	"analysis-question-AI/internal/api/cli"
)

type analysisCommand struct {
	flags 	*cli.flagsConfig
	svc 	*analysisService
	log 	*core.Logger
}


func newAnalysisCommand(log *core.Logger, flags *cli.flagsConfig, svc *analysisService) *analysisCommand {
	return &analysisCommand{
		log: log,
		flags: flags,
		svc: svc,
	}
}

func (a *analysisCommand) AnalyzeAndCorrectQuestions(questions []model.QuestionTable) {
	


	for _, q := range questions {
        analyzed, changed, err := a.svc.AnalyzeQuestions(questions)
        if err != nil { return nil, err }

        if changed {
            row := q.StartRow 
            if err := a.svc.CorrectQuestions(questions); err != nil {
                a.log.Printf("Ошибка обновления '%s'!A%d:E%d: %v", q.SheetName, row, row+3, err)
            }

			a.log.Info("Обновлен вопрос в листе ", q.SheetName, " Блок с строкой ", row, " - ", row+3)
        } 

        results = append(results, *analyzed)
    }
}