package analysis

import (
	"analysis-question-AI/internal/api/cli"
)

type AnalysisModule struct {
	command *analysisCommand
}


func NewAnalysisModule(log *core.Logger, flags *cli.flagsConfig) *AnalysisModule {

	api := newAnalysisApi();
	repo := newAnalysisRepository(api);
	svc := newAnalysisService(repo);
	command := newAnalysisCommand(log, flags, svc);

	return &AnalysisModule{
		command: command,
	}
}


func (a *AnalysisModule) Commands() *analysisCommand {
	return a.command; 
}