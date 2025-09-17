package sheet 

import (
	"analysis-question-AI/internal/api/cli"
	"analysis-question-AI/internal/core"
)

type SheetModule struct {
	commands *sheetCommand
}

func NewSheetModule(log *core.Logger, flags *cli.FlagsConfig, cfg *core.Config) *SheetModule {
	api := newSheetApi(cfg);
	repo := newSheetRepository(api);
	svc := newSheetService(repo, log);
	commands := newSheetCommand(log, flags, svc)

	return &SheetModule{
		commands: commands,
	}
}

func (s *SheetModule) Commands() *sheetCommand {
	return s.commands;
}
