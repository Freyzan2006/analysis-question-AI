package app

import (
	"analysis-question-AI/internal/modules/analysis"
	"analysis-question-AI/internal/modules/sheet"
	"analysis-question-AI/internal/api/cli"
	"analysis-question-AI/internal/core"
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
		sheetModule:    sheet.NewSheetModule(log, finalFlags, cfg),

		appLog:         log,
	}
}

func (c *CheckingCorrectingQuestionApplication) Run() {
	c.appLog.Info("Start checking and correcting questions")

	analysisCmd := c.analysisModule.Commands();
	sheetCmd := c.sheetModule.Commands();

	c.appLog.Info("Get all questions from sheet")
	allQuestion := sheetCmd.AllQuestions();

	c.appLog.Info("Analyze questions")
	correctedAnswers := analysisCmd.AnalyzeQuestions(allQuestion);
	

	c.appLog.Info("Correct questions")
	sheetCmd.SaveQuestions(correctedAnswers);

}