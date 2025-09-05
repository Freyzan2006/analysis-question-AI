package analysis

import (
	"analysis-question-AI/internal/api/cli"
)

type AnalysisModule struct {
	command *analysisCommand
}


func NewAnalysisModule(log *core.Logger, flags *cli.flagsConfig) *AnalysisModule {

	api := newAnalysisApi(log, flags.APIKey, flags.Model, flags.PromptTemplate);
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