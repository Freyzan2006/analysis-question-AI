package analysis

import (
	"analysis-question-AI/internal/api/cli"
	"analysis-question-AI/internal/core"
)

type AnalysisModule struct {
	command *analysisCommand
}


func NewAnalysisModule(log *core.Logger, flags *cli.FlagsConfig, env *core.Environment, cfg *core.Config) *AnalysisModule {
	
	API_GEMINI_KEY := env.Get("API_GEMINI_KEY")

	promptTemplate, err := core.LoadPrompt(cfg.PromptsPath)
	if err != nil {
		log.Fatal(err)
	}

	api := newAnalysisApi(log, API_GEMINI_KEY, "gemini-2.5-flash", promptTemplate);
	repo := newAnalysisRepository(api);
	svc := newAnalysisService(repo, log);
	command := newAnalysisCommand(svc, log, flags);

	return &AnalysisModule{
		command: command,
	}
}


func (a *AnalysisModule) Commands() *analysisCommand {
	return a.command; 
}