package app

import (
	"analysis-question-AI/internal/modules/analysis"
	"analysis-question-AI/internal/modules/sheet"
	"analysis-question-AI/internal/api/cli"
)

type CheckingCorrectingQuestionApplication struct {
	analysisModule 		*analysis.AnalysisModule
	sheetModule    		*sheet.SheetModule

	appLog          	*core.Logger
}

func NewCheckingCorrectingQuestionApplication(log *core.Logger) *CheckingCorrectingQuestionApplication {
	flags := cli.NewFlags()
	finalFlags := flags.GetFlags() 

	env := core.NewEnvironment()
	
	log := core.NewLogger(finalFlags.LogPathFile)
	cfg := &core.Config{
		SpreadsheetID:      finalFlags.GoogleSpreadsheetID,
		ReadRange:          finalFlags.GoogleReadRange,
		ServiceAccountFile: finalFlags.GoogleServiceAccountFile,
		PromptsPath:        finalFlags.GooglePromptsPath,
		Limit:              finalFlags.GoogleDocsLimit,
		Sheets:             finalFlags.GoogleDocsSheets,
	}


	return &CheckingCorrectingQuestionApplication{
		analysisModule: analysis.NewAnalysisModule(log, finalFlags, env, cfg),
		sheetModule:    analysis.NewSheetModule(log, finalFlags, cfg),

		appLog:         log,
	}
}

func (c *CheckingCorrectingQuestionApplication) Run() {
	appLog.Info("Start checking and correcting questions")

	analysisCmd := a.analysisModule.Commands();
	sheetCmd := a.sheetModule.Commands();

	appLog.Info("Get all questions from sheet")
	allQuestion := sheetCmd.AllQuestions();

	appLog.Info("Analyze questions")
	correctedAnswers := analysisCmd.AnalyzeQuestions(allQuestion);

	appLog.Info("Correct questions")
	sheetCmd.SaveQuestions(correctedAnswers);

}