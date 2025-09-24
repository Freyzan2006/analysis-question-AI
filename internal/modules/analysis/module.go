package analysis

import (
	"analysis-question-AI/internal/core"
)

type AnalysisModule struct {
	command *analysisCommand
}


func NewAnalysisModule(log *core.Logger, cfg *core.Config) *AnalysisModule {
	
	

	promptTemplate := getPromptFromFileUtil(cfg.PromptsPath);

	api := newAnalysisApi(log, cfg.ApiGeminiKey, "gemini-2.5-flash", promptTemplate);
	repo := newAnalysisRepository(api);
	svc := newAnalysisService(repo, log);
	command := newAnalysisCommand(svc, log, cfg);

	return &AnalysisModule{
		command: command,
	}
}


func (a *AnalysisModule) Commands() *analysisCommand {
	return a.command; 
}