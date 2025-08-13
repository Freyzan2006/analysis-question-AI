package main

import (
	"analysis-question-AI/internal/api/cli"
	"analysis-question-AI/internal/repository"
	"analysis-question-AI/internal/service"
	"analysis-question-AI/internal/core"
	"analysis-question-AI/internal/api/http/external"
)

import (
	"log"
)



// func main() {
// 	// инициализация переменных окружения
// 	var env = core.NewEnvironment()

// 	// инициализация конфига
// 	cfg, err := core.LoadConfig("config.json")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// инициализация api gemini
// 	var API_GEMINI_URL = env.Get("API_GEMINI_URL")
// 	var API_GEMINI_KEY = env.Get("API_GEMINI_KEY")
// 	var apiGemini = external.NewGeminiAPI(API_GEMINI_URL, API_GEMINI_KEY)


// 	// инициализация api google docs
// 	var API_GOOGLE_DOCS_URL = env.Get("API_GOOGLE_DOCS_URL")
// 	var apiGoogleDocs = external.NewGoogleDocsAPI(API_GOOGLE_DOCS_URL, cfg)

// 	// инициализация флагов из командного ряда
// 	var flags = cli.NewFlags()
	
// 	// Инициализация репозитория
// 	var repo = repository.NewQuestionRepository()

// 	// Инициализация сервисов
// 	var geminiSvc = service.NewQuestionService(apiGemini, repo)
// 	var googleDocsSvc = service.NewGoogleDocsService(apiGoogleDocs, repo)
// 	googleDocsSvc.GetQuestions()

// 	// Инициализация cli команд
// 	var commands = cli.NewCommands(flags, geminiSvc, repo, env)
// 	commands.Run()

// 	// test_features.Test()
// }


func main() {
	// инициализация переменных окружения
	env := core.NewEnvironment()

	// инициализация флагов (внутри они грузят config.json и подменяют значениями из CLI)
	flags := cli.NewFlags()
	finalFlags := flags.GetFlags() // <-- теперь это твой финальный конфиг
	cfg := &core.Config{
		SpreadsheetID:      finalFlags.GoogleSpreadsheetID,
		ReadRange:          finalFlags.GoogleReadRange,
		ServiceAccountFile: finalFlags.GoogleServiceAccountFile,
		PromptsPath:        finalFlags.GooglePromptsPath,
	}
	promptTemplate, err := core.LoadPrompt(cfg.PromptsPath)
	if err != nil {
		log.Fatal(err)
	}

	// инициализация API Gemini
	API_GEMINI_URL := env.Get("API_GEMINI_URL")
	API_GEMINI_KEY := env.Get("API_GEMINI_KEY")
	apiGemini := external.NewGeminiAPI(API_GEMINI_URL, API_GEMINI_KEY, promptTemplate)

	// инициализация API Google Docs (передаём финальный конфиг)
	API_GOOGLE_DOCS_URL := env.Get("API_GOOGLE_DOCS_URL")
	apiGoogleDocs := external.NewGoogleDocsAPI(API_GOOGLE_DOCS_URL, cfg)

	// инициализация репозитория
	repo := repository.NewQuestionRepository()

	// инициализация сервисов
	googleDocsSvc := service.NewGoogleDocsService(apiGoogleDocs, repo)
	geminiSvc := service.NewQuestionService(apiGemini, repo, googleDocsSvc)
	
	

	// инициализация CLI команд
	commands := cli.NewCommands(flags, geminiSvc, repo, env)
	commands.Run()
}
