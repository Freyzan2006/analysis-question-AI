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

func NewCheckingCorrectingQuestionApplication() *CheckingCorrectingQuestionApplication {
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
		LogPathFile:        finalFlags.LogPathFile,
		ApiGeminiKey:       env.Get("API_GEMINI_KEY"),
	}

	log := core.NewLogger(finalFlags.LogPathFile)


	return &CheckingCorrectingQuestionApplication{
		analysisModule: analysis.NewAnalysisModule(log, cfg),
		sheetModule:    sheet.NewSheetModule(log, cfg),

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
	

	c.appLog.Info("Save... corrected questions to sheet")
	sheetCmd.SaveQuestions(correctedAnswers);

}