package app

import (
	"analysis-question-AI/internal/modules/analysis"
	"analysis-question-AI/internal/api/cli"
)

type CheckingCorrectingQuestionApplication struct {
	analysisModule 		*analysis.AnalysisModule
	sheetModule    		*analysis.SheetModule

	appLog          	*core.Logger
}

func NewCheckingCorrectingQuestionApplication(log *core.Logger) *CheckingCorrectingQuestionApplication {
	flags := cli.NewFlags()
	flags = flags.GetFlags()

	return &CheckingCorrectingQuestionApplication{
		analysisModule: analysis.NewAnalysisModule(log, flags),
		sheetModule:    analysis.NewSheetModule(log, flags),

		appLog:         log,
	}
}

func (c *CheckingCorrectingQuestionApplication) Run() {
	appLog.Info("Start checking and correcting questions")

	analysisCmd := a.analysisModule.Commands();
	sheetCmd := a.sheetModule.Commands();

	appLog.Info("Get all questions from sheet")
	allQuestion := sheetCmd.AllQuestions();

	appLog.Info("Analyze and correct questions")
	analysisCmd.AnalyzeAndCorrectQuestions(allQuestion);

	appLog.Info("Get all changed questions")
	analysisCmd.AllChangedQuestions();
}