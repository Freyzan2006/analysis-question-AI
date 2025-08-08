package main

import (
	"analysis-question-AI/internal/api/cli"
	"analysis-question-AI/internal/repository"
	"analysis-question-AI/internal/service"
	"analysis-question-AI/internal/core"
	"analysis-question-AI/internal/api/http/external"
)


func main() {
	// инициализация переменных окружения
	var env = core.NewEnvironment()

	// инициализация api gemini
	var API_GEMINI_URL = env.Get("API_GEMINI_URL")
	var API_GEMINI_KEY = env.Get("API_GEMINI_KEY")
	var apiGemini = external.NewGeminiAPI(API_GEMINI_URL, API_GEMINI_KEY)


	// инициализация api google docs
	var API_GOOGLE_DOCS_URL = env.Get("API_GOOGLE_DOCS_URL")
	var apiGoogleDocs = external.NewGoogleDocsAPI(API_GOOGLE_DOCS_URL)

	// инициализация флагов из командного ряда
	var flags = cli.NewFlags()
	
	// Инициализация репозитория
	var repo = repository.NewQuestionRepository()

	// Инициализация сервисов
	var geminiSvc = service.NewQuestionService(apiGemini, repo)
	var googleDocsSvc = service.NewGoogleDocsService(apiGoogleDocs, repo)
	googleDocsSvc.GetQuestions()

	// Инициализация cli команд
	var commands = cli.NewCommands(flags, geminiSvc, repo, env)
	commands.Run()
}