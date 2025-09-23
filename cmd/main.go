package main

// import (
// 	"analysis-question-AI/internal/api/cli"
// 	"analysis-question-AI/internal/repository"
// 	"analysis-question-AI/internal/service"
// 	"analysis-question-AI/internal/core"
// 	"analysis-question-AI/internal/api/http/external"
// )

import (
	"analysis-question-AI/internal/app"
)



func main() {

	// // инициализация переменных окружения
	// env := core.NewEnvironment()

	// // инициализация флагов (внутри они грузят config.json и подменяют значениями из CLI)
	// flags := cli.NewFlags()
	// finalFlags := flags.GetFlags() 
	
	// log := core.NewLogger(finalFlags.LogPathFile)
	// cfg := &core.Config{
	// 	SpreadsheetID:      finalFlags.GoogleSpreadsheetID,
	// 	ReadRange:          finalFlags.GoogleReadRange,
	// 	ServiceAccountFile: finalFlags.GoogleServiceAccountFile,
	// 	PromptsPath:        finalFlags.GooglePromptsPath,
	// 	Limit:              finalFlags.GoogleDocsLimit,
	// 	Sheets:             finalFlags.GoogleDocsSheets,
	// }

	// promptTemplate, err := core.LoadPrompt(cfg.PromptsPath)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // инициализация API Gemini
	// // API_GEMINI_URL := env.Get("API_GEMINI_URL")
	// API_GEMINI_KEY := env.Get("API_GEMINI_KEY")
	// apiGemini := external.NewGeminiAPI(API_GEMINI_KEY, "gemini-2.5-flash", promptTemplate)

	// // инициализация API Google Docs (передаём финальный конфиг)
	// apiGoogleDocs := external.NewGoogleDocsAPI(cfg)

	// // инициализация репозитория
	// repo := repository.NewQuestionRepository(cfg)

	// // инициализация сервисов
	// googleDocsSvc := service.NewGoogleDocsService(apiGoogleDocs, repo)
	// questionSvc := service.NewQuestionService(apiGemini, repo, googleDocsSvc, log)
	
	

	// // инициализация CLI команд
	// commands := cli.NewCommands(flags, questionSvc, repo)
	// commands.Run()



	
	application := app.NewCheckingCorrectingQuestionApplication()
	application.Run()
}
